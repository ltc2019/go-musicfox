// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package backgroundtransfer

type BackgroundTransferCostPolicy int32

const SignatureBackgroundTransferCostPolicy string = "enum(Windows.Networking.BackgroundTransfer.BackgroundTransferCostPolicy;i4)"

const (
	BackgroundTransferCostPolicyDefault          BackgroundTransferCostPolicy = 0
	BackgroundTransferCostPolicyUnrestrictedOnly BackgroundTransferCostPolicy = 1
	BackgroundTransferCostPolicyAlways           BackgroundTransferCostPolicy = 2
)
