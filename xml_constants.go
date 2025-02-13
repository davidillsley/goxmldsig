package dsig

import "crypto"

const (
	DefaultPrefix = "ds"
	Namespace     = "http://www.w3.org/2000/09/xmldsig#"
)

// Tags
const (
	SignatureTag              = "Signature"
	SignedInfoTag             = "SignedInfo"
	CanonicalizationMethodTag = "CanonicalizationMethod"
	SignatureMethodTag        = "SignatureMethod"
	ReferenceTag              = "Reference"
	TransformsTag             = "Transforms"
	TransformTag              = "Transform"
	DigestMethodTag           = "DigestMethod"
	DigestValueTag            = "DigestValue"
	SignatureValueTag         = "SignatureValue"
	KeyInfoTag                = "KeyInfo"
	X509DataTag               = "X509Data"
	X509CertificateTag        = "X509Certificate"
)

const (
	AlgorithmAttr = "Algorithm"
	URIAttr       = "URI"
	DefaultIdAttr = "ID"
)

const (
	// NOTE(russell_h): I guess 1.0 is "exclusive" and 1.1 isn't
	CanonicalXML10AlgorithmId     = "http://www.w3.org/2001/10/xml-exc-c14n#"
	CanonicalXML11AlgorithmId     = "http://www.w3.org/2006/12/xml-c14n11"
	EnvelopedSignatureAltorithmId = "http://www.w3.org/2000/09/xmldsig#enveloped-signature"
)

var digestAlgorithmIdentifiers = map[crypto.Hash]string{
	crypto.SHA1:   "http://www.w3.org/2000/09/xmldsig#sha1",
	crypto.SHA256: "http://www.w3.org/2001/04/xmlenc#sha256",
}

var digestAlgorithmsByIdentifier = map[string]crypto.Hash{}
var signatureMethodsByIdentifier = map[string]crypto.Hash{}

func init() {
	for hash, id := range digestAlgorithmIdentifiers {
		digestAlgorithmsByIdentifier[id] = hash
	}
	for hash, id := range signatureMethodIdentifiers {
		signatureMethodsByIdentifier[id] = hash
	}
}

var signatureMethodIdentifiers = map[crypto.Hash]string{
	crypto.SHA1:   "http://www.w3.org/2000/09/xmldsig#rsa-sha1",
	crypto.SHA256: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
}
