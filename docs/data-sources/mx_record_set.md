---
page_title: "dns_mx_record_set Data Source - terraform-provider-dns"
subcategory: ""
description: |-
  Use this data source to get DNS MX records for a domain.
---

# dns_mx_record_set (Data Source)

Use this data source to get DNS MX records for a domain.

## Example Usage

```terraform
data "dns_mx_record_set" "mail" {
  domain = "example.com."
}

output "mailserver" {
  value = data.dns_mx_record_set.mail.mx.0.exchange
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String) Domain to look up.

### Read-Only

- `id` (String) Always set to the domain.
- `mx` (List of Object) A list of records. They are sorted by ascending preference then alphabetically by exchange to stay consistent across runs. (see [below for nested schema](#nestedatt--mx))

<a id="nestedatt--mx"></a>
### Nested Schema for `mx`

Read-Only:

- `exchange` (String)
- `preference` (Number)