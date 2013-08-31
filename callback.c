#include "callback.h"
#include "stdlib.h"

typedef struct {
	int setOffset;
	int getOffset;
} OffsetInfo;

void TW_CALL SetVar(void *value, void *clientData) {
	OffsetInfo* offsetInfo = (OffsetInfo*)clientData;
	goSetVarCallbackN(value, offsetInfo->setOffset);
}

void TW_CALL GetVar(void *value, void *clientData) {
	OffsetInfo* offsetInfo = (OffsetInfo*)clientData;
	goGetVarCallbackN(value, offsetInfo->getOffset);
}

void TW_CALL Button(void *clientData) {
	int offset = *((int*)clientData);
	goButtonCallbackN(offset);
}

int myAddVarCB(TwBar* bar, char* name, TwType type, int currentSetVar, int currentGetVar, char* def) {
	// this is a memory leak it will never be freed
	OffsetInfo* offsetInfo = (OffsetInfo*)malloc(sizeof(OffsetInfo));
	offsetInfo->setOffset = currentSetVar;
	offsetInfo->getOffset = currentGetVar;
	return TwAddVarCB(bar, name, type, (TwSetVarCallback)SetVar, GetVar, offsetInfo, def);
}

int myAddButton(TwBar *bar, char *name, int currentButton, char *def) {
	// this is a memory leak it will never be freed
	int *offset = (int*)malloc(sizeof(int));
	*offset = currentButton;
	return TwAddButton(bar, name, Button, offset, def);
}

void TW_CALL HandleError(const char *errorMessage){
	goHandleError((char*)errorMessage);
}

void myHandleErrors() {
	TwHandleErrors(HandleError);
}

void TW_CALL HandleSummary(char *summaryString, size_t summaryMaxLength, const void *value, void *clientData) {
	goSummary(summaryString, summaryMaxLength, (void*)value);
}

TwType myDefineStruct(const char *name, const TwStructMember *structMembers, unsigned int nbMembers, size_t structSize) {
	return TwDefineStruct(name, structMembers, nbMembers, structSize, HandleSummary, NULL);
}
