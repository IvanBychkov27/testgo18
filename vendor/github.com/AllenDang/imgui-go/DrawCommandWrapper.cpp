#include "DrawCommandWrapper.h"
#include "WrapperConverter.h"
#include "imguiWrappedHeader.h"

unsigned int iggDrawCommandIdxOffset(IggDrawCmd handle) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  return cmd->IdxOffset;
}

void iggDrawCommandGetVertexOffset(IggDrawCmd handle, unsigned int *count) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  *count = cmd->VtxOffset;
}

void iggDrawCommandGetElementCount(IggDrawCmd handle, unsigned int *count) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  *count = cmd->ElemCount;
}

void iggDrawCommandGetClipRect(IggDrawCmd handle, IggVec4 *rect) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  exportValue(*rect, cmd->ClipRect);
}

void iggDrawCommandGetTextureID(IggDrawCmd handle, IggTextureID *id) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  *id = reinterpret_cast<IggTextureID>(cmd->TextureId);
}

IggBool iggDrawCommandHasUserCallback(IggDrawCmd handle) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  return (cmd->UserCallback != 0) ? 1 : 0;
}

void iggDrawCommandCallUserCallback(IggDrawCmd handle, IggDrawList listHandle) {
  ImDrawCmd *cmd = reinterpret_cast<ImDrawCmd *>(handle);
  ImDrawList *list = reinterpret_cast<ImDrawList *>(listHandle);
  cmd->UserCallback(list, cmd);
}
