import { h } from "vue"
import { NTooltip, NButton } from "naive-ui";

const accessLevels = {
  account: {
    NONE: ["None", "error", undefined, "How did you get here??? Please, report this immideately"],
    READ: ["Read", "error", undefined, "You can only see this Account"],
    MGMT: ["Manage", "warning", undefined, "You can Manage this Account, for example enable/disable it"],
    ADMIN: ["Admin", "success", undefined, "You have the highest possible access to this Account"],
    ROOT: ["Super-Admin", "success", "#8a2be2", "You have the highest possible access to this Account"],
    OWNER: ["Owned", "success", "#8a2be2", "You are the owner of this Account, which gives you full access to it and right to delete it"]
  },
  namespace: {
    NONE: ["None", "error", undefined, "How did you get here??? Please, report this immideately"],
    READ: ["Read", "error", undefined, "You can only see this Namespace"],
    MGMT: ["Manage", "warning", undefined, "You can Manage this Namespace, for example enable/disable it"],
    ADMIN: ["Admin", "success", undefined, "You have the highest possible access to this Namespace"],
    ROOT: ["Super-Admin", "success", "#8a2be2", "You have the highest possible access to this Namespace"],
    OWNER: ["Owned", "success", "#8a2be2", "You are the owner of this Namespace, which gives you full access to it and right to delete it"]
  }
};

export default function AccessBadge(props) {
  let key = "account"
  if (props.account != undefined) key = "account"
  if (props.namespace != undefined) key = "namespace"

  let conf = accessLevels[key][props.access];
  return h(
    NTooltip,
    {
      trigger: "hover",
      placement: "top",
    },
    {
      trigger: () => h(
        NButton,
        {
          secondary: true,
          round: true,
          type: conf[1],
          color: conf[2],
          style: {
            marginLeft: props.left
          }
        },
        {
          default: () => conf[0],
        }
      ),
      default: () => conf[3]
    }
  )
}