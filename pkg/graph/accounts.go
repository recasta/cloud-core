/*
Copyright © 2021-2022 Infinite Devices GmbH, Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package graph

import (
	"context"

	"github.com/arangodb/go-driver"
	"github.com/golang-jwt/jwt"
	"github.com/infinimesh/infinimesh/pkg/credentials"
	"github.com/infinimesh/infinimesh/pkg/graph/schema"
	pb "github.com/infinimesh/infinimesh/pkg/node/proto"
	"github.com/infinimesh/infinimesh/pkg/node/proto/access"
	accpb "github.com/infinimesh/infinimesh/pkg/node/proto/accounts"
	inf "github.com/infinimesh/infinimesh/pkg/shared"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Account struct {
	*accpb.Account
	driver.DocumentMeta
}

func (o *Account) ID() (driver.DocumentID) {
	return o.DocumentMeta.ID
}

func (o *Account) SetAccessLevel(level access.AccessLevel) {
	if o.Access == nil {
		o.Access = &access.Access{
			Level: level,
		}
		return
	}
	o.Access.Level = level
}

func NewBlankAccountDocument(key string) *Account {
	return &Account{
		Account: &accpb.Account{
			Uuid: key,
		},
		DocumentMeta: NewBlankDocument(schema.ACCOUNTS_COL, key),
	}
}

func NewAccountFromPB(acc *accpb.Account) (res *Account) {
	return &Account{
		Account: acc,
		DocumentMeta: NewBlankDocument(schema.ACCOUNTS_COL, acc.Uuid),
	}
}

type AccountsController struct {
	pb.UnimplementedAccountsServiceServer
	log *zap.Logger

	col driver.Collection // Accounts Collection
	cred driver.Collection
	db driver.Database

	acc2ns driver.Collection // Accounts to Namespaces permissions edge collection
	ns2acc driver.Collection // Namespaces to Accounts permissions edge collection

	SIGNING_KEY []byte
}

func NewAccountsController(log *zap.Logger, db driver.Database) *AccountsController {
	ctx := context.TODO()
	perm_graph, _ := db.Graph(ctx, schema.PERMISSIONS_GRAPH.Name)
	col, _ := perm_graph.VertexCollection(ctx, schema.ACCOUNTS_COL)

	cred_graph, _ := db.Graph(ctx, schema.CREDENTIALS_GRAPH.Name)
	cred, _ := cred_graph.VertexCollection(ctx, schema.CREDENTIALS_COL)
	return &AccountsController{
		log: log.Named("AccountsController"), col: col, db: db, cred: cred,
		acc2ns: GetEdgeCol(ctx, db, schema.ACC2NS), ns2acc: GetEdgeCol(ctx, db, schema.NS2ACC),
		SIGNING_KEY: []byte("just-an-init-thing-replace-me"),
	}
}

func (c *AccountsController) Token(ctx context.Context, req *pb.TokenRequest) (*pb.TokenResponse, error) {
	log := c.log.Named("Token")
	log.Debug("Token request received", zap.Any("request", req))

	account, ok := c.Authorize(ctx, req.Auth.Type, req.Auth.Data...)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Wrong credentials given")
	}
	log.Debug("Authorized user", zap.String("ID", account.ID().String()))
	if !account.Enabled {
		return nil, status.Error(codes.PermissionDenied, "Account is disabled")
	}

	claims := jwt.MapClaims{}
	claims[inf.INFINIMESH_ACCOUNT_CLAIM] = account.Key
	claims["exp"] = req.Exp

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(c.SIGNING_KEY)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to issue token")
	}

	return &pb.TokenResponse{Token: token_string}, nil
}

func (c *AccountsController) Get(ctx context.Context, acc *accpb.Account) (res *accpb.Account, err error) {
	log := c.log.Named("Get")
	log.Debug("Get request received", zap.Any("request", acc))

	requestor := ctx.Value(inf.InfinimeshAccountCtxKey).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	uuid := acc.GetUuid()
	if uuid == "me" {
		uuid = requestor
	}
	// Getting Account from DB
	// and Check requestor access
	result := *NewBlankAccountDocument(uuid)
	err = AccessLevelAndGet(ctx, log, c.db, NewBlankAccountDocument(requestor), &result)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Account not found or not enough Access Rights")
	}
	if result.Access.Level < access.AccessLevel_READ {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access Rights")
	}

	return result.Account, nil
}

func (c *AccountsController) List(ctx context.Context, _ *pb.EmptyMessage) (*accpb.Accounts, error) {
	log := c.log.Named("List")

	requestor := ctx.Value(inf.InfinimeshAccountCtxKey).(string)
	log.Debug("Requestor", zap.String("id", requestor))


	cr, err := ListQuery(ctx, log, c.db, NewBlankAccountDocument(requestor), schema.ACCOUNTS_COL, 4)
	if err != nil {
		log.Error("Error executing query", zap.Error(err))
		return nil, status.Error(codes.Internal, "Couldn't execute query")
	}
	defer cr.Close()

	var r []*accpb.Account
	for {
		var acc accpb.Account
		meta, err := cr.ReadDocument(ctx, &acc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			log.Error("Error unmarshalling Document", zap.Error(err))
			return nil, status.Error(codes.Internal, "Couldn't execute query")
		}
		acc.Uuid = meta.ID.Key()
		log.Debug("Got document", zap.Any("account", &acc))
		r = append(r, &acc)
	}

	return &accpb.Accounts{
		Accounts: r,
	}, nil
}

func (c *AccountsController) Create(ctx context.Context, request *accpb.CreateRequest) (*accpb.CreateResponse, error) {
	log := c.log.Named("Create")
	log.Debug("Create request received", zap.Any("request", request), zap.Any("context", ctx))

	requestor := ctx.Value(inf.InfinimeshAccountCtxKey).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	ns_id := request.GetNamespace()
	if ns_id == "" {
		ns_id = schema.ROOT_NAMESPACE_KEY
	}

	ok, level := AccessLevel(ctx, c.db, NewBlankAccountDocument(requestor), NewBlankNamespaceDocument(ns_id))
	if !ok || level < access.AccessLevel_ADMIN {
		return nil, status.Errorf(codes.PermissionDenied, "No Access to Namespace %s", ns_id)
	}

	if request.Account.GetDefaultNamespace() == "" {
		request.Account.DefaultNamespace = ns_id
	}

	account := Account{Account: request.GetAccount()}
	meta, err := c.col.CreateDocument(ctx, account)
	if err != nil {
		log.Error("Error creating Account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while creating Account")
	}
	account.Uuid = meta.ID.Key()
	account.DocumentMeta = meta

	ns := NewBlankNamespaceDocument(ns_id)
	err = Link(ctx, log, c.ns2acc, ns, &account, access.AccessLevel_ADMIN, access.Role_UNSET)
	if err != nil {
		defer c.col.RemoveDocument(ctx, meta.Key)
		log.Error("Error Linking Namespace to Account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	col, _ := c.db.Collection(ctx, schema.CREDENTIALS_EDGE_COL)
	cred, err := credentials.MakeCredentials(request.GetCredentials(), log)
	if err != nil {
		defer c.col.RemoveDocument(ctx, meta.Key)
		log.Error("Error making Credentials for Account", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = c.SetCredentialsCtrl(ctx, account, col, cred)
	if err != nil {
		defer c.col.RemoveDocument(ctx, meta.Key)
		log.Error("Error setting Credentials for Account", zap.Error(err))
		return nil, err
	}

	return &accpb.CreateResponse{Account: account.Account}, nil
}

func (c *AccountsController) Update(ctx context.Context, acc *accpb.Account) (*accpb.Account, error) {
	log := c.log.Named("Update")
	log.Debug("Update request received", zap.Any("request", acc), zap.Any("context", ctx))

	requestor := ctx.Value(inf.InfinimeshAccountCtxKey).(string)
	log.Debug("Requestor", zap.String("id", requestor))
	requestorAccount := NewBlankAccountDocument(requestor)	

	old := *NewBlankAccountDocument(acc.GetUuid())
	err := AccessLevelAndGet(ctx, log, c.db, requestorAccount, &old)
	if err != nil || old.Access.Level < access.AccessLevel_ADMIN {
		return nil, status.Errorf(codes.PermissionDenied, "No Access to Account %s", acc.GetUuid())
	}

	if old.GetDefaultNamespace() != acc.GetDefaultNamespace() {
		ok, level := AccessLevel(ctx, c.db, requestorAccount, NewBlankNamespaceDocument(acc.GetDefaultNamespace()))
		if !ok || level < access.AccessLevel_READ {
			return nil, status.Errorf(codes.PermissionDenied, "No Access to Namespace %s", acc.GetDefaultNamespace())
		}
	}

	_, err = c.col.UpdateDocument(ctx, acc.GetUuid(), acc)
	if err != nil {
		log.Error("Internal error while updating Document", zap.Any("request", acc), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating Account")
	}

	return acc, nil
}

func (c *AccountsController) Toggle(ctx context.Context, acc *accpb.Account) (*accpb.Account, error) {
	log := c.log.Named("Update")
	log.Debug("Update request received", zap.Any("account", acc), zap.Any("context", ctx))

	curr, err := c.Get(ctx, acc)
	if err != nil {
		return nil, err
	}

	if curr.Access.Level < access.AccessLevel_MGMT {
		return nil, status.Errorf(codes.PermissionDenied, "No Access to Account %s", acc.Uuid)
	}

	res := NewAccountFromPB(curr)
	err = Toggle(ctx, c.db, res, "enabled")
	if err != nil {
		log.Error("Error updating Account", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error while updating Account")
	}

	return res.Account, nil
}

func (c *AccountsController) Delete(ctx context.Context, req *accpb.Account) (*pb.DeleteResponse, error)  {
	log := c.log.Named("Delete")
	log.Debug("Delete request received", zap.Any("request", req), zap.Any("context", ctx))

	requestor := ctx.Value(inf.InfinimeshAccountCtxKey).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	acc := *NewBlankAccountDocument(req.GetUuid())
	err := AccessLevelAndGet(ctx, log, c.db, NewBlankAccountDocument(requestor), &acc)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Account not found or not enough Access Rights")
	}
	if acc.Access.Level < access.AccessLevel_ADMIN {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access Rights")
	}

	creds, err := c.GetCredentials(ctx, acc)
	if err != nil {
		log.Error("Error gathering Account credentials", zap.String("account", acc.Key), zap.Error(err))
	}
	log.Debug("Got credentials", zap.Any("credentials", creds))

	_, errs, err := c.cred.RemoveDocuments(ctx, creds)
	if err != nil {
		log.Error("Error deleting Credentials", zap.String("account", acc.Key), zap.Any("errors", errs), zap.Error(err))
		return nil, status.Error(codes.Internal, "Account has been deleted partialy")
	}

	_, err = c.col.RemoveDocument(ctx, acc.ID().Key())
	if err != nil {
		log.Error("Error deleting Account", zap.String("account", acc.Key), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error deleting Account")
	}

	return &pb.DeleteResponse{}, nil
}

// Helper Functions

func (ctrl *AccountsController) Authorize(ctx context.Context, auth_type string, args ...string) (Account, bool) {
	ctrl.log.Debug("Authorization request", zap.String("type", auth_type))

	credentials, err := credentials.Find(ctx, ctrl.col.Database(), ctrl.log, auth_type, args...)
	// Check if could authorize
	if err != nil {
		ctrl.log.Info("Coudn't authorize", zap.Error(err))
		return Account{}, false
	}

	account, ok := Authorisable(ctx, &credentials, ctrl.col.Database())
	ctrl.log.Debug("Authorized account", zap.Bool("result", ok), zap.Any("account", account))
	return account, ok
}

// Return Account authorisable by this Credentials
func Authorisable(ctx context.Context, cred *credentials.Credentials, db driver.Database) (Account, bool) {
	query := `FOR account IN 1 INBOUND @credentials GRAPH @credentials_graph RETURN account`
	c, err := db.Query(ctx, query, map[string]interface{}{
		"credentials": cred,
		"credentials_graph": schema.CREDENTIALS_GRAPH.Name,
	})
	if err != nil {
		return Account{}, false
	}
	defer c.Close()

	var r Account
	_, err = c.ReadDocument(ctx, &r)
	return r, err == nil
}

// Return Credentials linked to Account
func (ctrl *AccountsController) GetCredentials(ctx context.Context, acc Account) (r []string, err error) {
	query := `FOR credentials IN 1 OUTBOUND @account GRAPH @credentials_graph RETURN credentials._key`
	c, err := ctrl.db.Query(ctx, query, map[string]interface{}{
		"account": acc.ID().String(),
		"credentials_graph": schema.CREDENTIALS_GRAPH.Name,
	})
	if err != nil {
		ctrl.log.Error("Error executing query", zap.Error(err))
		return nil, err
	}
	defer c.Close()

	for {
		var cred string
		_, err := c.ReadDocument(ctx, &cred)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			ctrl.log.Debug("Error unmarshalling credentials", zap.Error(err))
			return nil, err
		}
		r = append(r, cred)
	}
	return r, nil
}

// Set Account Credentials, ensure account has only one credentials document linked per credentials type
func (ctrl *AccountsController) SetCredentialsCtrl(ctx context.Context, acc Account, edge driver.Collection, c credentials.Credentials) (error) {
	key := c.Type() + "-" + acc.Key
	var oldLink credentials.Link
	meta, err := edge.ReadDocument(ctx, key, &oldLink)
	if err == nil {	
		ctrl.log.Debug("Link exists", zap.Any("meta", meta))
		_, err = ctrl.cred.UpdateDocument(ctx, oldLink.To.Key(), c)
		if err != nil {
			ctrl.log.Error("Error updating Credentials of type", zap.Error(err), zap.String("key", key))
			return status.Error(codes.InvalidArgument, "Error updating Credentials of type")
		}

		return nil
	}
	ctrl.log.Debug("Credentials either not created yet or failed to get them from DB, overwriting", zap.Error(err), zap.String("key", key))

	cred, err := ctrl.cred.CreateDocument(ctx, c)	
	if err != nil {
		ctrl.log.Error("Error creating Credentials Document", zap.String("type", c.Type()), zap.Error(err))
		return status.Error(codes.Internal, "Couldn't create credentials")
	}

	_, err = edge.CreateDocument(ctx, credentials.Link{
		From: acc.ID(),
		To: cred.ID,
		Type: c.Type(),
		DocumentMeta: driver.DocumentMeta {
			Key: key, // Ensures only one credentials vertex per type
		},
	})
	if err != nil {
		ctrl.log.Error("Error Linking Credentials to Account",
			zap.String("account", acc.Key), zap.String("type", c.Type()), zap.Error(err),
		)
		ctrl.cred.RemoveDocument(ctx, cred.Key)
		return status.Error(codes.Internal, "Couldn't assign credentials")
	}
	return nil
}

func (c *AccountsController) SetCredentials(ctx context.Context, req *pb.SetCredentialsRequest) (*pb.SetCredentialsResponse, error) {
	log := c.log.Named("SetCredentials")
	log.Debug("Set Credentials request received", zap.String("account", req.GetUuid()), zap.String("type", req.GetCredentials().GetType()), zap.Any("context", ctx))

	requestor := ctx.Value(inf.InfinimeshAccountCtxKey).(string)
	log.Debug("Requestor", zap.String("id", requestor))

	acc := *NewBlankAccountDocument(req.GetUuid())
	err := AccessLevelAndGet(ctx, log, c.db, NewBlankAccountDocument(requestor), &acc)

	if err != nil {
		log.Error("Error getting Account", zap.String("requestor", requestor), zap.String("account", req.GetUuid()) , zap.Error(err))
		return nil, status.Error(codes.Internal, "Error getting Account or not enough Access right to set credentials for this Account")
	}

	if acc.Access.Level < access.AccessLevel_ROOT || acc.Access.Role != access.Role_OWNER {
		return nil, status.Error(codes.PermissionDenied, "Not enough Access right to set credentials for this Account. Only Owner and Super-Admin can do this")
	}

	col, _ := c.db.Collection(ctx, schema.CREDENTIALS_EDGE_COL)
	cred, err := credentials.MakeCredentials(req.GetCredentials(), log)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = c.SetCredentialsCtrl(ctx, acc, col, cred)
	if err != nil {
		return nil, err
	}
	return &pb.SetCredentialsResponse{}, nil
}
