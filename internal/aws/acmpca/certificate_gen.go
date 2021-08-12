// Code generated by generators/resource/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_acmpca_certificate", certificateResourceType)
}

// certificateResourceType returns the Terraform aws_acmpca_certificate resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ACMPCA::Certificate resource type.
func certificateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"api_passthrough": {
			// Property: ApiPassthrough
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Structure that specifies fields to be overridden in a certificate at the time of issuance. These requires an API Passthrough template be used or they will be ignored.",
			//   "properties": {
			//     "Extensions": {
			//       "additionalProperties": false,
			//       "description": "Structure that contains X.500 extensions for a Certificate.",
			//       "properties": {
			//         "CertificatePolicies": {
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Structure that contains X.509 Policy information.",
			//             "properties": {
			//               "CertPolicyId": {
			//                 "description": "String that contains X.509 ObjectIdentifier information.",
			//                 "type": "string"
			//               },
			//               "PolicyQualifiers": {
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "description": "Structure that contains X.509 Policy qualifier information.",
			//                   "properties": {
			//                     "PolicyQualifierId": {
			//                       "type": "string"
			//                     },
			//                     "Qualifier": {
			//                       "additionalProperties": false,
			//                       "description": "Structure that contains a X.509 policy qualifier.",
			//                       "properties": {
			//                         "CpsUri": {
			//                           "type": "string"
			//                         }
			//                       },
			//                       "required": [
			//                         "CpsUri"
			//                       ],
			//                       "type": "object"
			//                     }
			//                   },
			//                   "required": [
			//                     "PolicyQualifierId",
			//                     "Qualifier"
			//                   ],
			//                   "type": "object"
			//                 },
			//                 "type": "array"
			//               }
			//             },
			//             "required": [
			//               "CertPolicyId"
			//             ],
			//             "type": "object"
			//           },
			//           "type": "array"
			//         },
			//         "ExtendedKeyUsage": {
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Structure that contains X.509 ExtendedKeyUsage information.",
			//             "properties": {
			//               "ExtendedKeyUsageObjectIdentifier": {
			//                 "description": "String that contains X.509 ObjectIdentifier information.",
			//                 "type": "string"
			//               },
			//               "ExtendedKeyUsageType": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         },
			//         "KeyUsage": {
			//           "additionalProperties": false,
			//           "description": "Structure that contains X.509 KeyUsage information.",
			//           "properties": {
			//             "CRLSign": {
			//               "type": "boolean"
			//             },
			//             "DataEncipherment": {
			//               "type": "boolean"
			//             },
			//             "DecipherOnly": {
			//               "type": "boolean"
			//             },
			//             "DigitalSignature": {
			//               "type": "boolean"
			//             },
			//             "EncipherOnly": {
			//               "type": "boolean"
			//             },
			//             "KeyAgreement": {
			//               "type": "boolean"
			//             },
			//             "KeyCertSign": {
			//               "type": "boolean"
			//             },
			//             "KeyEncipherment": {
			//               "type": "boolean"
			//             },
			//             "NonRepudiation": {
			//               "type": "boolean"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SubjectAlternativeNames": {
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "Structure that contains X.509 GeneralName information. Assign one and ONLY one field.",
			//             "properties": {
			//               "DirectoryName": {
			//                 "additionalProperties": false,
			//                 "description": "Structure that contains X.500 distinguished name information.",
			//                 "properties": {
			//                   "CommonName": {
			//                     "type": "string"
			//                   },
			//                   "Country": {
			//                     "type": "string"
			//                   },
			//                   "DistinguishedNameQualifier": {
			//                     "type": "string"
			//                   },
			//                   "GenerationQualifier": {
			//                     "type": "string"
			//                   },
			//                   "GivenName": {
			//                     "type": "string"
			//                   },
			//                   "Initials": {
			//                     "type": "string"
			//                   },
			//                   "Locality": {
			//                     "type": "string"
			//                   },
			//                   "Organization": {
			//                     "type": "string"
			//                   },
			//                   "OrganizationalUnit": {
			//                     "type": "string"
			//                   },
			//                   "Pseudonym": {
			//                     "type": "string"
			//                   },
			//                   "SerialNumber": {
			//                     "type": "string"
			//                   },
			//                   "State": {
			//                     "type": "string"
			//                   },
			//                   "Surname": {
			//                     "type": "string"
			//                   },
			//                   "Title": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "DnsName": {
			//                 "description": "String that contains X.509 DnsName information.",
			//                 "type": "string"
			//               },
			//               "EdiPartyName": {
			//                 "additionalProperties": false,
			//                 "description": "Structure that contains X.509 EdiPartyName information.",
			//                 "properties": {
			//                   "NameAssigner": {
			//                     "type": "string"
			//                   },
			//                   "PartyName": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "PartyName",
			//                   "NameAssigner"
			//                 ],
			//                 "type": "object"
			//               },
			//               "IpAddress": {
			//                 "description": "String that contains X.509 IpAddress information.",
			//                 "type": "string"
			//               },
			//               "OtherName": {
			//                 "additionalProperties": false,
			//                 "description": "Structure that contains X.509 OtherName information.",
			//                 "properties": {
			//                   "TypeId": {
			//                     "description": "String that contains X.509 ObjectIdentifier information.",
			//                     "type": "string"
			//                   },
			//                   "Value": {
			//                     "type": "string"
			//                   }
			//                 },
			//                 "required": [
			//                   "TypeId",
			//                   "Value"
			//                 ],
			//                 "type": "object"
			//               },
			//               "RegisteredId": {
			//                 "description": "String that contains X.509 ObjectIdentifier information.",
			//                 "type": "string"
			//               },
			//               "Rfc822Name": {
			//                 "description": "String that contains X.509 Rfc822Name information.",
			//                 "type": "string"
			//               },
			//               "UniformResourceIdentifier": {
			//                 "description": "String that contains X.509 UniformResourceIdentifier information.",
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "Subject": {
			//       "additionalProperties": false,
			//       "description": "Structure that contains X.500 distinguished name information.",
			//       "properties": {
			//         "CommonName": {
			//           "type": "string"
			//         },
			//         "Country": {
			//           "type": "string"
			//         },
			//         "DistinguishedNameQualifier": {
			//           "type": "string"
			//         },
			//         "GenerationQualifier": {
			//           "type": "string"
			//         },
			//         "GivenName": {
			//           "type": "string"
			//         },
			//         "Initials": {
			//           "type": "string"
			//         },
			//         "Locality": {
			//           "type": "string"
			//         },
			//         "Organization": {
			//           "type": "string"
			//         },
			//         "OrganizationalUnit": {
			//           "type": "string"
			//         },
			//         "Pseudonym": {
			//           "type": "string"
			//         },
			//         "SerialNumber": {
			//           "type": "string"
			//         },
			//         "State": {
			//           "type": "string"
			//         },
			//         "Surname": {
			//           "type": "string"
			//         },
			//         "Title": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Structure that specifies fields to be overridden in a certificate at the time of issuance. These requires an API Passthrough template be used or they will be ignored.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"extensions": {
						// Property: Extensions
						Description: "Structure that contains X.500 extensions for a Certificate.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"certificate_policies": {
									// Property: CertificatePolicies
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"cert_policy_id": {
												// Property: CertPolicyId
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Required:    true,
											},
											"policy_qualifiers": {
												// Property: PolicyQualifiers
												Attributes: schema.ListNestedAttributes(
													map[string]schema.Attribute{
														"policy_qualifier_id": {
															// Property: PolicyQualifierId
															Type:     types.StringType,
															Required: true,
														},
														"qualifier": {
															// Property: Qualifier
															Description: "Structure that contains a X.509 policy qualifier.",
															Attributes: schema.SingleNestedAttributes(
																map[string]schema.Attribute{
																	"cps_uri": {
																		// Property: CpsUri
																		Type:     types.StringType,
																		Required: true,
																	},
																},
															),
															Required: true,
														},
													},
													schema.ListNestedAttributesOptions{},
												),
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"extended_key_usage": {
									// Property: ExtendedKeyUsage
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"extended_key_usage_object_identifier": {
												// Property: ExtendedKeyUsageObjectIdentifier
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"extended_key_usage_type": {
												// Property: ExtendedKeyUsageType
												Type:     types.StringType,
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
								"key_usage": {
									// Property: KeyUsage
									Description: "Structure that contains X.509 KeyUsage information.",
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"crl_sign": {
												// Property: CRLSign
												Type:     types.BoolType,
												Optional: true,
											},
											"data_encipherment": {
												// Property: DataEncipherment
												Type:     types.BoolType,
												Optional: true,
											},
											"decipher_only": {
												// Property: DecipherOnly
												Type:     types.BoolType,
												Optional: true,
											},
											"digital_signature": {
												// Property: DigitalSignature
												Type:     types.BoolType,
												Optional: true,
											},
											"encipher_only": {
												// Property: EncipherOnly
												Type:     types.BoolType,
												Optional: true,
											},
											"key_agreement": {
												// Property: KeyAgreement
												Type:     types.BoolType,
												Optional: true,
											},
											"key_cert_sign": {
												// Property: KeyCertSign
												Type:     types.BoolType,
												Optional: true,
											},
											"key_encipherment": {
												// Property: KeyEncipherment
												Type:     types.BoolType,
												Optional: true,
											},
											"non_repudiation": {
												// Property: NonRepudiation
												Type:     types.BoolType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"subject_alternative_names": {
									// Property: SubjectAlternativeNames
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"directory_name": {
												// Property: DirectoryName
												Description: "Structure that contains X.500 distinguished name information.",
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"common_name": {
															// Property: CommonName
															Type:     types.StringType,
															Optional: true,
														},
														"country": {
															// Property: Country
															Type:     types.StringType,
															Optional: true,
														},
														"distinguished_name_qualifier": {
															// Property: DistinguishedNameQualifier
															Type:     types.StringType,
															Optional: true,
														},
														"generation_qualifier": {
															// Property: GenerationQualifier
															Type:     types.StringType,
															Optional: true,
														},
														"given_name": {
															// Property: GivenName
															Type:     types.StringType,
															Optional: true,
														},
														"initials": {
															// Property: Initials
															Type:     types.StringType,
															Optional: true,
														},
														"locality": {
															// Property: Locality
															Type:     types.StringType,
															Optional: true,
														},
														"organization": {
															// Property: Organization
															Type:     types.StringType,
															Optional: true,
														},
														"organizational_unit": {
															// Property: OrganizationalUnit
															Type:     types.StringType,
															Optional: true,
														},
														"pseudonym": {
															// Property: Pseudonym
															Type:     types.StringType,
															Optional: true,
														},
														"serial_number": {
															// Property: SerialNumber
															Type:     types.StringType,
															Optional: true,
														},
														"state": {
															// Property: State
															Type:     types.StringType,
															Optional: true,
														},
														"surname": {
															// Property: Surname
															Type:     types.StringType,
															Optional: true,
														},
														"title": {
															// Property: Title
															Type:     types.StringType,
															Optional: true,
														},
													},
												),
												Optional: true,
											},
											"dns_name": {
												// Property: DnsName
												Description: "String that contains X.509 DnsName information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"edi_party_name": {
												// Property: EdiPartyName
												Description: "Structure that contains X.509 EdiPartyName information.",
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"name_assigner": {
															// Property: NameAssigner
															Type:     types.StringType,
															Required: true,
														},
														"party_name": {
															// Property: PartyName
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"ip_address": {
												// Property: IpAddress
												Description: "String that contains X.509 IpAddress information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"other_name": {
												// Property: OtherName
												Description: "Structure that contains X.509 OtherName information.",
												Attributes: schema.SingleNestedAttributes(
													map[string]schema.Attribute{
														"type_id": {
															// Property: TypeId
															Description: "String that contains X.509 ObjectIdentifier information.",
															Type:        types.StringType,
															Required:    true,
														},
														"value": {
															// Property: Value
															Type:     types.StringType,
															Required: true,
														},
													},
												),
												Optional: true,
											},
											"registered_id": {
												// Property: RegisteredId
												Description: "String that contains X.509 ObjectIdentifier information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"rfc_822_name": {
												// Property: Rfc822Name
												Description: "String that contains X.509 Rfc822Name information.",
												Type:        types.StringType,
												Optional:    true,
											},
											"uniform_resource_identifier": {
												// Property: UniformResourceIdentifier
												Description: "String that contains X.509 UniformResourceIdentifier information.",
												Type:        types.StringType,
												Optional:    true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"subject": {
						// Property: Subject
						Description: "Structure that contains X.500 distinguished name information.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"common_name": {
									// Property: CommonName
									Type:     types.StringType,
									Optional: true,
								},
								"country": {
									// Property: Country
									Type:     types.StringType,
									Optional: true,
								},
								"distinguished_name_qualifier": {
									// Property: DistinguishedNameQualifier
									Type:     types.StringType,
									Optional: true,
								},
								"generation_qualifier": {
									// Property: GenerationQualifier
									Type:     types.StringType,
									Optional: true,
								},
								"given_name": {
									// Property: GivenName
									Type:     types.StringType,
									Optional: true,
								},
								"initials": {
									// Property: Initials
									Type:     types.StringType,
									Optional: true,
								},
								"locality": {
									// Property: Locality
									Type:     types.StringType,
									Optional: true,
								},
								"organization": {
									// Property: Organization
									Type:     types.StringType,
									Optional: true,
								},
								"organizational_unit": {
									// Property: OrganizationalUnit
									Type:     types.StringType,
									Optional: true,
								},
								"pseudonym": {
									// Property: Pseudonym
									Type:     types.StringType,
									Optional: true,
								},
								"serial_number": {
									// Property: SerialNumber
									Type:     types.StringType,
									Optional: true,
								},
								"state": {
									// Property: State
									Type:     types.StringType,
									Optional: true,
								},
								"surname": {
									// Property: Surname
									Type:     types.StringType,
									Optional: true,
								},
								"title": {
									// Property: Title
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// ApiPassthrough is a force-new attribute.
			// ApiPassthrough is a write-only attribute.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"certificate": {
			// Property: Certificate
			// CloudFormation resource type schema:
			// {
			//   "description": "The issued certificate in base 64 PEM-encoded format.",
			//   "type": "string"
			// }
			Description: "The issued certificate in base 64 PEM-encoded format.",
			Type:        types.StringType,
			Computed:    true,
			// Certificate is a force-new attribute.
			// Certificate is a write-only attribute.
		},
		"certificate_authority_arn": {
			// Property: CertificateAuthorityArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// CertificateAuthorityArn is a force-new attribute.
		},
		"certificate_signing_request": {
			// Property: CertificateSigningRequest
			// CloudFormation resource type schema:
			// {
			//   "description": "The certificate signing request (CSR) for the Certificate.",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The certificate signing request (CSR) for the Certificate.",
			Type:        types.StringType,
			Required:    true,
			// CertificateSigningRequest is a force-new attribute.
			// CertificateSigningRequest is a write-only attribute.
		},
		"signing_algorithm": {
			// Property: SigningAlgorithm
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the algorithm that will be used to sign the Certificate.",
			//   "type": "string"
			// }
			Description: "The name of the algorithm that will be used to sign the Certificate.",
			Type:        types.StringType,
			Required:    true,
			// SigningAlgorithm is a force-new attribute.
		},
		"template_arn": {
			// Property: TemplateArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// TemplateArn is a force-new attribute.
		},
		"validity": {
			// Property: Validity
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Validity for a certificate.",
			//   "properties": {
			//     "Type": {
			//       "type": "string"
			//     },
			//     "Value": {
			//       "type": "number"
			//     }
			//   },
			//   "required": [
			//     "Value",
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Description: "Validity for a certificate.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Required: true,
			// Validity is a force-new attribute.
		},
		"validity_not_before": {
			// Property: ValidityNotBefore
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Validity for a certificate.",
			//   "properties": {
			//     "Type": {
			//       "type": "string"
			//     },
			//     "Value": {
			//       "type": "number"
			//     }
			//   },
			//   "required": [
			//     "Value",
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Description: "Validity for a certificate.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.NumberType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// ValidityNotBefore is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "A certificate issued via a private certificate authority",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::Certificate").WithTerraformTypeName("aws_acmpca_certificate").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ApiPassthrough",
		"/properties/CertificateSigningRequest",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_acmpca_certificate", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
