package types

import (
	sdkerrors "cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

var ModuleName string

var (
	ConvertToString = "[Type Error] Conversion failed: cannot convert from Int to String."
	ParseUint       = "[Parse Error] Failed to parse the provided data. Please ensure the input format is correct."
	ErrEmptyAddress = "[Address Error] The provided address is empty. AccAddressFromBech32 requires a non-empty address string."
	InvalidRequest  = "[Request Error] The request is invalid. Please check the parameters and try again."
)

var (
	ErrInvalidRequest = sdkerrors.Register(ModuleName, 1, "Invalid Request: The request format is invalid or missing required parameters.")
	ErrNotFound       = sdkerrors.RegisterWithGRPCCode(ModuleName, 5, codes.NotFound, "not found")
)

var (
	ErrKeyNotFound = sdkerrors.Register(ModuleName, 2, "Key Not Found: The specified key could not be located in the data store.")
)

var (
	ErrInvalidType    = sdkerrors.Register(ModuleName, 3, "Invalid Type: The provided value is of an unexpected type and cannot be processed.")
	ErrUnknownRequest = sdkerrors.Register(ModuleName, 4, "Unknown Request: The request type is unrecognized. Please verify the request details.")
)
