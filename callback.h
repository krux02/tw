#ifndef CALLBACK_H
#define CALLBACK_H

#include <AntTweakBar.h>

extern void goSetVarCallbackN(void *value, int N);
extern void goGetVarCallbackN(void *value, int N);
extern void goButtonCallbackN(int N);
extern void goHandleError(char *errorMessage);
extern void goSummary(char *summaryString, size_t summaryMaxLength, void *value);

int myAddVarCB(TwBar* bar, char* name, TwType type, int currentSetVar, int currentGetVar, char* def);
int myAddButton(TwBar *bar, char *name, int currentButton, char *def);
void myHandleErrors();
TwStructMember myCreateStructMember(char* name, TwType type, size_t offset, char* defString);
TwType myDefineStruct(const char *name, const TwStructMember *structMembers, unsigned int nbMembers, size_t structSize);

#endif
