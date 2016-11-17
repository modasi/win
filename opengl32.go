// +build windows

package win

import (
	"unsafe"
)

var (
	// Library
	libopengl32 uintptr

	// Functions
	glAccum                   uintptr
	glAlphaFunc               uintptr
	glAreTexturesResident     uintptr
	glArrayElement            uintptr
	glBegin                   uintptr
	glBindTexture             uintptr
	glBitmap                  uintptr
	glBlendFunc               uintptr
	glCallList                uintptr
	glCallLists               uintptr
	glClear                   uintptr
	glClearAccum              uintptr
	glClearColor              uintptr
	glClearDepth              uintptr
	glClearIndex              uintptr
	glClearStencil            uintptr
	glClipPlane               uintptr
	glColor3b                 uintptr
	glColor3bv                uintptr
	glColor3d                 uintptr
	glColor3dv                uintptr
	glColor3f                 uintptr
	glColor3fv                uintptr
	glColor3i                 uintptr
	glColor3iv                uintptr
	glColor3s                 uintptr
	glColor3sv                uintptr
	glColor3ub                uintptr
	glColor3ubv               uintptr
	glColor3ui                uintptr
	glColor3uiv               uintptr
	glColor3us                uintptr
	glColor3usv               uintptr
	glColor4b                 uintptr
	glColor4bv                uintptr
	glColor4d                 uintptr
	glColor4dv                uintptr
	glColor4f                 uintptr
	glColor4fv                uintptr
	glColor4i                 uintptr
	glColor4iv                uintptr
	glColor4s                 uintptr
	glColor4sv                uintptr
	glColor4ub                uintptr
	glColor4ubv               uintptr
	glColor4ui                uintptr
	glColor4uiv               uintptr
	glColor4us                uintptr
	glColor4usv               uintptr
	glColorMask               uintptr
	glColorMaterial           uintptr
	glColorPointer            uintptr
	glCopyPixels              uintptr
	glCopyTexImage1D          uintptr
	glCopyTexImage2D          uintptr
	glCopyTexSubImage1D       uintptr
	glCopyTexSubImage2D       uintptr
	glCullFace                uintptr
	glDeleteLists             uintptr
	glDeleteTextures          uintptr
	glDepthFunc               uintptr
	glDepthMask               uintptr
	glDepthRange              uintptr
	glDisable                 uintptr
	glDisableClientState      uintptr
	glDrawArrays              uintptr
	glDrawBuffer              uintptr
	glDrawElements            uintptr
	glDrawPixels              uintptr
	glEdgeFlag                uintptr
	glEdgeFlagPointer         uintptr
	glEdgeFlagv               uintptr
	glEnable                  uintptr
	glEnableClientState       uintptr
	glEnd                     uintptr
	glEndList                 uintptr
	glEvalCoord1d             uintptr
	glEvalCoord1dv            uintptr
	glEvalCoord1f             uintptr
	glEvalCoord1fv            uintptr
	glEvalCoord2d             uintptr
	glEvalCoord2dv            uintptr
	glEvalCoord2f             uintptr
	glEvalCoord2fv            uintptr
	glEvalMesh1               uintptr
	glEvalMesh2               uintptr
	glEvalPoint1              uintptr
	glEvalPoint2              uintptr
	glFeedbackBuffer          uintptr
	glFinish                  uintptr
	glFlush                   uintptr
	glFogf                    uintptr
	glFogfv                   uintptr
	glFogi                    uintptr
	glFogiv                   uintptr
	glFrontFace               uintptr
	glFrustum                 uintptr
	glGenLists                uintptr
	glGenTextures             uintptr
	glGetBooleanv             uintptr
	glGetClipPlane            uintptr
	glGetDoublev              uintptr
	glGetError                uintptr
	glGetFloatv               uintptr
	glGetIntegerv             uintptr
	glGetLightfv              uintptr
	glGetLightiv              uintptr
	glGetMapdv                uintptr
	glGetMapfv                uintptr
	glGetMapiv                uintptr
	glGetMaterialfv           uintptr
	glGetMaterialiv           uintptr
	glGetPixelMapfv           uintptr
	glGetPixelMapuiv          uintptr
	glGetPixelMapusv          uintptr
	glGetPointerv             uintptr
	glGetPolygonStipple       uintptr
	glGetTexEnvfv             uintptr
	glGetTexEnviv             uintptr
	glGetTexGendv             uintptr
	glGetTexGenfv             uintptr
	glGetTexGeniv             uintptr
	glGetTexImage             uintptr
	glGetTexLevelParameterfv  uintptr
	glGetTexLevelParameteriv  uintptr
	glGetTexParameterfv       uintptr
	glGetTexParameteriv       uintptr
	glHint                    uintptr
	glIndexMask               uintptr
	glIndexPointer            uintptr
	glIndexd                  uintptr
	glIndexdv                 uintptr
	glIndexf                  uintptr
	glIndexfv                 uintptr
	glIndexi                  uintptr
	glIndexiv                 uintptr
	glIndexs                  uintptr
	glIndexsv                 uintptr
	glIndexub                 uintptr
	glIndexubv                uintptr
	glInitNames               uintptr
	glInterleavedArrays       uintptr
	glIsEnabled               uintptr
	glIsList                  uintptr
	glIsTexture               uintptr
	glLightModelf             uintptr
	glLightModelfv            uintptr
	glLightModeli             uintptr
	glLightModeliv            uintptr
	glLightf                  uintptr
	glLightfv                 uintptr
	glLighti                  uintptr
	glLightiv                 uintptr
	glLineStipple             uintptr
	glLineWidth               uintptr
	glListBase                uintptr
	glLoadIdentity            uintptr
	glLoadMatrixd             uintptr
	glLoadMatrixf             uintptr
	glLoadName                uintptr
	glLogicOp                 uintptr
	glMap1d                   uintptr
	glMap1f                   uintptr
	glMap2d                   uintptr
	glMap2f                   uintptr
	glMapGrid1d               uintptr
	glMapGrid1f               uintptr
	glMapGrid2d               uintptr
	glMapGrid2f               uintptr
	glMaterialf               uintptr
	glMaterialfv              uintptr
	glMateriali               uintptr
	glMaterialiv              uintptr
	glMatrixMode              uintptr
	glMultMatrixd             uintptr
	glMultMatrixf             uintptr
	glNewList                 uintptr
	glNormal3b                uintptr
	glNormal3bv               uintptr
	glNormal3d                uintptr
	glNormal3dv               uintptr
	glNormal3f                uintptr
	glNormal3fv               uintptr
	glNormal3i                uintptr
	glNormal3iv               uintptr
	glNormal3s                uintptr
	glNormal3sv               uintptr
	glNormalPointer           uintptr
	glOrtho                   uintptr
	glPassThrough             uintptr
	glPixelMapfv              uintptr
	glPixelMapuiv             uintptr
	glPixelMapusv             uintptr
	glPixelStoref             uintptr
	glPixelStorei             uintptr
	glPixelTransferf          uintptr
	glPixelTransferi          uintptr
	glPixelZoom               uintptr
	glPointSize               uintptr
	glPolygonMode             uintptr
	glPolygonOffset           uintptr
	glPolygonStipple          uintptr
	glPopAttrib               uintptr
	glPopClientAttrib         uintptr
	glPopMatrix               uintptr
	glPopName                 uintptr
	glPrioritizeTextures      uintptr
	glPushAttrib              uintptr
	glPushClientAttrib        uintptr
	glPushMatrix              uintptr
	glPushName                uintptr
	glRasterPos2d             uintptr
	glRasterPos2dv            uintptr
	glRasterPos2f             uintptr
	glRasterPos2fv            uintptr
	glRasterPos2i             uintptr
	glRasterPos2iv            uintptr
	glRasterPos2s             uintptr
	glRasterPos2sv            uintptr
	glRasterPos3d             uintptr
	glRasterPos3dv            uintptr
	glRasterPos3f             uintptr
	glRasterPos3fv            uintptr
	glRasterPos3i             uintptr
	glRasterPos3iv            uintptr
	glRasterPos3s             uintptr
	glRasterPos3sv            uintptr
	glRasterPos4d             uintptr
	glRasterPos4dv            uintptr
	glRasterPos4f             uintptr
	glRasterPos4fv            uintptr
	glRasterPos4i             uintptr
	glRasterPos4iv            uintptr
	glRasterPos4s             uintptr
	glRasterPos4sv            uintptr
	glReadBuffer              uintptr
	glReadPixels              uintptr
	glRectd                   uintptr
	glRectdv                  uintptr
	glRectf                   uintptr
	glRectfv                  uintptr
	glRecti                   uintptr
	glRectiv                  uintptr
	glRects                   uintptr
	glRectsv                  uintptr
	glRenderMode              uintptr
	glRotated                 uintptr
	glRotatef                 uintptr
	glScaled                  uintptr
	glScalef                  uintptr
	glScissor                 uintptr
	glSelectBuffer            uintptr
	glShadeModel              uintptr
	glStencilFunc             uintptr
	glStencilMask             uintptr
	glStencilOp               uintptr
	glTexCoord1d              uintptr
	glTexCoord1dv             uintptr
	glTexCoord1f              uintptr
	glTexCoord1fv             uintptr
	glTexCoord1i              uintptr
	glTexCoord1iv             uintptr
	glTexCoord1s              uintptr
	glTexCoord1sv             uintptr
	glTexCoord2d              uintptr
	glTexCoord2dv             uintptr
	glTexCoord2f              uintptr
	glTexCoord2fv             uintptr
	glTexCoord2i              uintptr
	glTexCoord2iv             uintptr
	glTexCoord2s              uintptr
	glTexCoord2sv             uintptr
	glTexCoord3d              uintptr
	glTexCoord3dv             uintptr
	glTexCoord3f              uintptr
	glTexCoord3fv             uintptr
	glTexCoord3i              uintptr
	glTexCoord3iv             uintptr
	glTexCoord3s              uintptr
	glTexCoord3sv             uintptr
	glTexCoord4d              uintptr
	glTexCoord4dv             uintptr
	glTexCoord4f              uintptr
	glTexCoord4fv             uintptr
	glTexCoord4i              uintptr
	glTexCoord4iv             uintptr
	glTexCoord4s              uintptr
	glTexCoord4sv             uintptr
	glTexCoordPointer         uintptr
	glTexEnvf                 uintptr
	glTexEnvfv                uintptr
	glTexEnvi                 uintptr
	glTexEnviv                uintptr
	glTexGend                 uintptr
	glTexGendv                uintptr
	glTexGenf                 uintptr
	glTexGenfv                uintptr
	glTexGeni                 uintptr
	glTexGeniv                uintptr
	glTexImage1D              uintptr
	glTexImage2D              uintptr
	glTexParameterf           uintptr
	glTexParameterfv          uintptr
	glTexParameteri           uintptr
	glTexParameteriv          uintptr
	glTexSubImage1D           uintptr
	glTexSubImage2D           uintptr
	glTranslated              uintptr
	glTranslatef              uintptr
	glVertex2d                uintptr
	glVertex2dv               uintptr
	glVertex2f                uintptr
	glVertex2fv               uintptr
	glVertex2i                uintptr
	glVertex2iv               uintptr
	glVertex2s                uintptr
	glVertex2sv               uintptr
	glVertex3d                uintptr
	glVertex3dv               uintptr
	glVertex3f                uintptr
	glVertex3fv               uintptr
	glVertex3i                uintptr
	glVertex3iv               uintptr
	glVertex3s                uintptr
	glVertex3sv               uintptr
	glVertex4d                uintptr
	glVertex4dv               uintptr
	glVertex4f                uintptr
	glVertex4fv               uintptr
	glVertex4i                uintptr
	glVertex4iv               uintptr
	glVertex4s                uintptr
	glVertex4sv               uintptr
	glVertexPointer           uintptr
	glViewport                uintptr
	wglCopyContext            uintptr
	wglCreateContext          uintptr
	wglCreateLayerContext     uintptr
	wglDeleteContext          uintptr
	wglDescribeLayerPlane     uintptr
	wglGetCurrentContext      uintptr
	wglGetCurrentDC           uintptr
	wglGetLayerPaletteEntries uintptr
	wglGetProcAddress         uintptr
	wglMakeCurrent            uintptr
	wglRealizeLayerPalette    uintptr
	wglSetLayerPaletteEntries uintptr
	wglShareLists             uintptr
	wglSwapLayerBuffers       uintptr
	wglSwapMultipleBuffers    uintptr
	wglUseFontBitmaps         uintptr
	wglUseFontOutlines        uintptr
	glDebugEntry              uintptr
	wglChoosePixelFormat      uintptr
	wglDescribePixelFormat    uintptr
	wglGetPixelFormat         uintptr
	wglSetPixelFormat         uintptr
)

func init() {
	// Library
	libopengl32 = doLoadLibrary("opengl32.dll")

	// Functions
	glAccum = doGetProcAddress(libopengl32, "glAccum")
	glAlphaFunc = doGetProcAddress(libopengl32, "glAlphaFunc")
	glAreTexturesResident = doGetProcAddress(libopengl32, "glAreTexturesResident")
	glArrayElement = doGetProcAddress(libopengl32, "glArrayElement")
	glBegin = doGetProcAddress(libopengl32, "glBegin")
	glBindTexture = doGetProcAddress(libopengl32, "glBindTexture")
	glBitmap = doGetProcAddress(libopengl32, "glBitmap")
	glBlendFunc = doGetProcAddress(libopengl32, "glBlendFunc")
	glCallList = doGetProcAddress(libopengl32, "glCallList")
	glCallLists = doGetProcAddress(libopengl32, "glCallLists")
	glClear = doGetProcAddress(libopengl32, "glClear")
	glClearAccum = doGetProcAddress(libopengl32, "glClearAccum")
	glClearColor = doGetProcAddress(libopengl32, "glClearColor")
	glClearDepth = doGetProcAddress(libopengl32, "glClearDepth")
	glClearIndex = doGetProcAddress(libopengl32, "glClearIndex")
	glClearStencil = doGetProcAddress(libopengl32, "glClearStencil")
	glClipPlane = doGetProcAddress(libopengl32, "glClipPlane")
	glColor3b = doGetProcAddress(libopengl32, "glColor3b")
	glColor3bv = doGetProcAddress(libopengl32, "glColor3bv")
	glColor3d = doGetProcAddress(libopengl32, "glColor3d")
	glColor3dv = doGetProcAddress(libopengl32, "glColor3dv")
	glColor3f = doGetProcAddress(libopengl32, "glColor3f")
	glColor3fv = doGetProcAddress(libopengl32, "glColor3fv")
	glColor3i = doGetProcAddress(libopengl32, "glColor3i")
	glColor3iv = doGetProcAddress(libopengl32, "glColor3iv")
	glColor3s = doGetProcAddress(libopengl32, "glColor3s")
	glColor3sv = doGetProcAddress(libopengl32, "glColor3sv")
	glColor3ub = doGetProcAddress(libopengl32, "glColor3ub")
	glColor3ubv = doGetProcAddress(libopengl32, "glColor3ubv")
	glColor3ui = doGetProcAddress(libopengl32, "glColor3ui")
	glColor3uiv = doGetProcAddress(libopengl32, "glColor3uiv")
	glColor3us = doGetProcAddress(libopengl32, "glColor3us")
	glColor3usv = doGetProcAddress(libopengl32, "glColor3usv")
	glColor4b = doGetProcAddress(libopengl32, "glColor4b")
	glColor4bv = doGetProcAddress(libopengl32, "glColor4bv")
	glColor4d = doGetProcAddress(libopengl32, "glColor4d")
	glColor4dv = doGetProcAddress(libopengl32, "glColor4dv")
	glColor4f = doGetProcAddress(libopengl32, "glColor4f")
	glColor4fv = doGetProcAddress(libopengl32, "glColor4fv")
	glColor4i = doGetProcAddress(libopengl32, "glColor4i")
	glColor4iv = doGetProcAddress(libopengl32, "glColor4iv")
	glColor4s = doGetProcAddress(libopengl32, "glColor4s")
	glColor4sv = doGetProcAddress(libopengl32, "glColor4sv")
	glColor4ub = doGetProcAddress(libopengl32, "glColor4ub")
	glColor4ubv = doGetProcAddress(libopengl32, "glColor4ubv")
	glColor4ui = doGetProcAddress(libopengl32, "glColor4ui")
	glColor4uiv = doGetProcAddress(libopengl32, "glColor4uiv")
	glColor4us = doGetProcAddress(libopengl32, "glColor4us")
	glColor4usv = doGetProcAddress(libopengl32, "glColor4usv")
	glColorMask = doGetProcAddress(libopengl32, "glColorMask")
	glColorMaterial = doGetProcAddress(libopengl32, "glColorMaterial")
	glColorPointer = doGetProcAddress(libopengl32, "glColorPointer")
	glCopyPixels = doGetProcAddress(libopengl32, "glCopyPixels")
	glCopyTexImage1D = doGetProcAddress(libopengl32, "glCopyTexImage1D")
	glCopyTexImage2D = doGetProcAddress(libopengl32, "glCopyTexImage2D")
	glCopyTexSubImage1D = doGetProcAddress(libopengl32, "glCopyTexSubImage1D")
	glCopyTexSubImage2D = doGetProcAddress(libopengl32, "glCopyTexSubImage2D")
	glCullFace = doGetProcAddress(libopengl32, "glCullFace")
	glDeleteLists = doGetProcAddress(libopengl32, "glDeleteLists")
	glDeleteTextures = doGetProcAddress(libopengl32, "glDeleteTextures")
	glDepthFunc = doGetProcAddress(libopengl32, "glDepthFunc")
	glDepthMask = doGetProcAddress(libopengl32, "glDepthMask")
	glDepthRange = doGetProcAddress(libopengl32, "glDepthRange")
	glDisable = doGetProcAddress(libopengl32, "glDisable")
	glDisableClientState = doGetProcAddress(libopengl32, "glDisableClientState")
	glDrawArrays = doGetProcAddress(libopengl32, "glDrawArrays")
	glDrawBuffer = doGetProcAddress(libopengl32, "glDrawBuffer")
	glDrawElements = doGetProcAddress(libopengl32, "glDrawElements")
	glDrawPixels = doGetProcAddress(libopengl32, "glDrawPixels")
	glEdgeFlag = doGetProcAddress(libopengl32, "glEdgeFlag")
	glEdgeFlagPointer = doGetProcAddress(libopengl32, "glEdgeFlagPointer")
	glEdgeFlagv = doGetProcAddress(libopengl32, "glEdgeFlagv")
	glEnable = doGetProcAddress(libopengl32, "glEnable")
	glEnableClientState = doGetProcAddress(libopengl32, "glEnableClientState")
	glEnd = doGetProcAddress(libopengl32, "glEnd")
	glEndList = doGetProcAddress(libopengl32, "glEndList")
	glEvalCoord1d = doGetProcAddress(libopengl32, "glEvalCoord1d")
	glEvalCoord1dv = doGetProcAddress(libopengl32, "glEvalCoord1dv")
	glEvalCoord1f = doGetProcAddress(libopengl32, "glEvalCoord1f")
	glEvalCoord1fv = doGetProcAddress(libopengl32, "glEvalCoord1fv")
	glEvalCoord2d = doGetProcAddress(libopengl32, "glEvalCoord2d")
	glEvalCoord2dv = doGetProcAddress(libopengl32, "glEvalCoord2dv")
	glEvalCoord2f = doGetProcAddress(libopengl32, "glEvalCoord2f")
	glEvalCoord2fv = doGetProcAddress(libopengl32, "glEvalCoord2fv")
	glEvalMesh1 = doGetProcAddress(libopengl32, "glEvalMesh1")
	glEvalMesh2 = doGetProcAddress(libopengl32, "glEvalMesh2")
	glEvalPoint1 = doGetProcAddress(libopengl32, "glEvalPoint1")
	glEvalPoint2 = doGetProcAddress(libopengl32, "glEvalPoint2")
	glFeedbackBuffer = doGetProcAddress(libopengl32, "glFeedbackBuffer")
	glFinish = doGetProcAddress(libopengl32, "glFinish")
	glFlush = doGetProcAddress(libopengl32, "glFlush")
	glFogf = doGetProcAddress(libopengl32, "glFogf")
	glFogfv = doGetProcAddress(libopengl32, "glFogfv")
	glFogi = doGetProcAddress(libopengl32, "glFogi")
	glFogiv = doGetProcAddress(libopengl32, "glFogiv")
	glFrontFace = doGetProcAddress(libopengl32, "glFrontFace")
	glFrustum = doGetProcAddress(libopengl32, "glFrustum")
	glGenLists = doGetProcAddress(libopengl32, "glGenLists")
	glGenTextures = doGetProcAddress(libopengl32, "glGenTextures")
	glGetBooleanv = doGetProcAddress(libopengl32, "glGetBooleanv")
	glGetClipPlane = doGetProcAddress(libopengl32, "glGetClipPlane")
	glGetDoublev = doGetProcAddress(libopengl32, "glGetDoublev")
	glGetError = doGetProcAddress(libopengl32, "glGetError")
	glGetFloatv = doGetProcAddress(libopengl32, "glGetFloatv")
	glGetIntegerv = doGetProcAddress(libopengl32, "glGetIntegerv")
	glGetLightfv = doGetProcAddress(libopengl32, "glGetLightfv")
	glGetLightiv = doGetProcAddress(libopengl32, "glGetLightiv")
	glGetMapdv = doGetProcAddress(libopengl32, "glGetMapdv")
	glGetMapfv = doGetProcAddress(libopengl32, "glGetMapfv")
	glGetMapiv = doGetProcAddress(libopengl32, "glGetMapiv")
	glGetMaterialfv = doGetProcAddress(libopengl32, "glGetMaterialfv")
	glGetMaterialiv = doGetProcAddress(libopengl32, "glGetMaterialiv")
	glGetPixelMapfv = doGetProcAddress(libopengl32, "glGetPixelMapfv")
	glGetPixelMapuiv = doGetProcAddress(libopengl32, "glGetPixelMapuiv")
	glGetPixelMapusv = doGetProcAddress(libopengl32, "glGetPixelMapusv")
	glGetPointerv = doGetProcAddress(libopengl32, "glGetPointerv")
	glGetPolygonStipple = doGetProcAddress(libopengl32, "glGetPolygonStipple")
	glGetTexEnvfv = doGetProcAddress(libopengl32, "glGetTexEnvfv")
	glGetTexEnviv = doGetProcAddress(libopengl32, "glGetTexEnviv")
	glGetTexGendv = doGetProcAddress(libopengl32, "glGetTexGendv")
	glGetTexGenfv = doGetProcAddress(libopengl32, "glGetTexGenfv")
	glGetTexGeniv = doGetProcAddress(libopengl32, "glGetTexGeniv")
	glGetTexImage = doGetProcAddress(libopengl32, "glGetTexImage")
	glGetTexLevelParameterfv = doGetProcAddress(libopengl32, "glGetTexLevelParameterfv")
	glGetTexLevelParameteriv = doGetProcAddress(libopengl32, "glGetTexLevelParameteriv")
	glGetTexParameterfv = doGetProcAddress(libopengl32, "glGetTexParameterfv")
	glGetTexParameteriv = doGetProcAddress(libopengl32, "glGetTexParameteriv")
	glHint = doGetProcAddress(libopengl32, "glHint")
	glIndexMask = doGetProcAddress(libopengl32, "glIndexMask")
	glIndexPointer = doGetProcAddress(libopengl32, "glIndexPointer")
	glIndexd = doGetProcAddress(libopengl32, "glIndexd")
	glIndexdv = doGetProcAddress(libopengl32, "glIndexdv")
	glIndexf = doGetProcAddress(libopengl32, "glIndexf")
	glIndexfv = doGetProcAddress(libopengl32, "glIndexfv")
	glIndexi = doGetProcAddress(libopengl32, "glIndexi")
	glIndexiv = doGetProcAddress(libopengl32, "glIndexiv")
	glIndexs = doGetProcAddress(libopengl32, "glIndexs")
	glIndexsv = doGetProcAddress(libopengl32, "glIndexsv")
	glIndexub = doGetProcAddress(libopengl32, "glIndexub")
	glIndexubv = doGetProcAddress(libopengl32, "glIndexubv")
	glInitNames = doGetProcAddress(libopengl32, "glInitNames")
	glInterleavedArrays = doGetProcAddress(libopengl32, "glInterleavedArrays")
	glIsEnabled = doGetProcAddress(libopengl32, "glIsEnabled")
	glIsList = doGetProcAddress(libopengl32, "glIsList")
	glIsTexture = doGetProcAddress(libopengl32, "glIsTexture")
	glLightModelf = doGetProcAddress(libopengl32, "glLightModelf")
	glLightModelfv = doGetProcAddress(libopengl32, "glLightModelfv")
	glLightModeli = doGetProcAddress(libopengl32, "glLightModeli")
	glLightModeliv = doGetProcAddress(libopengl32, "glLightModeliv")
	glLightf = doGetProcAddress(libopengl32, "glLightf")
	glLightfv = doGetProcAddress(libopengl32, "glLightfv")
	glLighti = doGetProcAddress(libopengl32, "glLighti")
	glLightiv = doGetProcAddress(libopengl32, "glLightiv")
	glLineStipple = doGetProcAddress(libopengl32, "glLineStipple")
	glLineWidth = doGetProcAddress(libopengl32, "glLineWidth")
	glListBase = doGetProcAddress(libopengl32, "glListBase")
	glLoadIdentity = doGetProcAddress(libopengl32, "glLoadIdentity")
	glLoadMatrixd = doGetProcAddress(libopengl32, "glLoadMatrixd")
	glLoadMatrixf = doGetProcAddress(libopengl32, "glLoadMatrixf")
	glLoadName = doGetProcAddress(libopengl32, "glLoadName")
	glLogicOp = doGetProcAddress(libopengl32, "glLogicOp")
	glMap1d = doGetProcAddress(libopengl32, "glMap1d")
	glMap1f = doGetProcAddress(libopengl32, "glMap1f")
	glMap2d = doGetProcAddress(libopengl32, "glMap2d")
	glMap2f = doGetProcAddress(libopengl32, "glMap2f")
	glMapGrid1d = doGetProcAddress(libopengl32, "glMapGrid1d")
	glMapGrid1f = doGetProcAddress(libopengl32, "glMapGrid1f")
	glMapGrid2d = doGetProcAddress(libopengl32, "glMapGrid2d")
	glMapGrid2f = doGetProcAddress(libopengl32, "glMapGrid2f")
	glMaterialf = doGetProcAddress(libopengl32, "glMaterialf")
	glMaterialfv = doGetProcAddress(libopengl32, "glMaterialfv")
	glMateriali = doGetProcAddress(libopengl32, "glMateriali")
	glMaterialiv = doGetProcAddress(libopengl32, "glMaterialiv")
	glMatrixMode = doGetProcAddress(libopengl32, "glMatrixMode")
	glMultMatrixd = doGetProcAddress(libopengl32, "glMultMatrixd")
	glMultMatrixf = doGetProcAddress(libopengl32, "glMultMatrixf")
	glNewList = doGetProcAddress(libopengl32, "glNewList")
	glNormal3b = doGetProcAddress(libopengl32, "glNormal3b")
	glNormal3bv = doGetProcAddress(libopengl32, "glNormal3bv")
	glNormal3d = doGetProcAddress(libopengl32, "glNormal3d")
	glNormal3dv = doGetProcAddress(libopengl32, "glNormal3dv")
	glNormal3f = doGetProcAddress(libopengl32, "glNormal3f")
	glNormal3fv = doGetProcAddress(libopengl32, "glNormal3fv")
	glNormal3i = doGetProcAddress(libopengl32, "glNormal3i")
	glNormal3iv = doGetProcAddress(libopengl32, "glNormal3iv")
	glNormal3s = doGetProcAddress(libopengl32, "glNormal3s")
	glNormal3sv = doGetProcAddress(libopengl32, "glNormal3sv")
	glNormalPointer = doGetProcAddress(libopengl32, "glNormalPointer")
	glOrtho = doGetProcAddress(libopengl32, "glOrtho")
	glPassThrough = doGetProcAddress(libopengl32, "glPassThrough")
	glPixelMapfv = doGetProcAddress(libopengl32, "glPixelMapfv")
	glPixelMapuiv = doGetProcAddress(libopengl32, "glPixelMapuiv")
	glPixelMapusv = doGetProcAddress(libopengl32, "glPixelMapusv")
	glPixelStoref = doGetProcAddress(libopengl32, "glPixelStoref")
	glPixelStorei = doGetProcAddress(libopengl32, "glPixelStorei")
	glPixelTransferf = doGetProcAddress(libopengl32, "glPixelTransferf")
	glPixelTransferi = doGetProcAddress(libopengl32, "glPixelTransferi")
	glPixelZoom = doGetProcAddress(libopengl32, "glPixelZoom")
	glPointSize = doGetProcAddress(libopengl32, "glPointSize")
	glPolygonMode = doGetProcAddress(libopengl32, "glPolygonMode")
	glPolygonOffset = doGetProcAddress(libopengl32, "glPolygonOffset")
	glPolygonStipple = doGetProcAddress(libopengl32, "glPolygonStipple")
	glPopAttrib = doGetProcAddress(libopengl32, "glPopAttrib")
	glPopClientAttrib = doGetProcAddress(libopengl32, "glPopClientAttrib")
	glPopMatrix = doGetProcAddress(libopengl32, "glPopMatrix")
	glPopName = doGetProcAddress(libopengl32, "glPopName")
	glPrioritizeTextures = doGetProcAddress(libopengl32, "glPrioritizeTextures")
	glPushAttrib = doGetProcAddress(libopengl32, "glPushAttrib")
	glPushClientAttrib = doGetProcAddress(libopengl32, "glPushClientAttrib")
	glPushMatrix = doGetProcAddress(libopengl32, "glPushMatrix")
	glPushName = doGetProcAddress(libopengl32, "glPushName")
	glRasterPos2d = doGetProcAddress(libopengl32, "glRasterPos2d")
	glRasterPos2dv = doGetProcAddress(libopengl32, "glRasterPos2dv")
	glRasterPos2f = doGetProcAddress(libopengl32, "glRasterPos2f")
	glRasterPos2fv = doGetProcAddress(libopengl32, "glRasterPos2fv")
	glRasterPos2i = doGetProcAddress(libopengl32, "glRasterPos2i")
	glRasterPos2iv = doGetProcAddress(libopengl32, "glRasterPos2iv")
	glRasterPos2s = doGetProcAddress(libopengl32, "glRasterPos2s")
	glRasterPos2sv = doGetProcAddress(libopengl32, "glRasterPos2sv")
	glRasterPos3d = doGetProcAddress(libopengl32, "glRasterPos3d")
	glRasterPos3dv = doGetProcAddress(libopengl32, "glRasterPos3dv")
	glRasterPos3f = doGetProcAddress(libopengl32, "glRasterPos3f")
	glRasterPos3fv = doGetProcAddress(libopengl32, "glRasterPos3fv")
	glRasterPos3i = doGetProcAddress(libopengl32, "glRasterPos3i")
	glRasterPos3iv = doGetProcAddress(libopengl32, "glRasterPos3iv")
	glRasterPos3s = doGetProcAddress(libopengl32, "glRasterPos3s")
	glRasterPos3sv = doGetProcAddress(libopengl32, "glRasterPos3sv")
	glRasterPos4d = doGetProcAddress(libopengl32, "glRasterPos4d")
	glRasterPos4dv = doGetProcAddress(libopengl32, "glRasterPos4dv")
	glRasterPos4f = doGetProcAddress(libopengl32, "glRasterPos4f")
	glRasterPos4fv = doGetProcAddress(libopengl32, "glRasterPos4fv")
	glRasterPos4i = doGetProcAddress(libopengl32, "glRasterPos4i")
	glRasterPos4iv = doGetProcAddress(libopengl32, "glRasterPos4iv")
	glRasterPos4s = doGetProcAddress(libopengl32, "glRasterPos4s")
	glRasterPos4sv = doGetProcAddress(libopengl32, "glRasterPos4sv")
	glReadBuffer = doGetProcAddress(libopengl32, "glReadBuffer")
	glReadPixels = doGetProcAddress(libopengl32, "glReadPixels")
	glRectd = doGetProcAddress(libopengl32, "glRectd")
	glRectdv = doGetProcAddress(libopengl32, "glRectdv")
	glRectf = doGetProcAddress(libopengl32, "glRectf")
	glRectfv = doGetProcAddress(libopengl32, "glRectfv")
	glRecti = doGetProcAddress(libopengl32, "glRecti")
	glRectiv = doGetProcAddress(libopengl32, "glRectiv")
	glRects = doGetProcAddress(libopengl32, "glRects")
	glRectsv = doGetProcAddress(libopengl32, "glRectsv")
	glRenderMode = doGetProcAddress(libopengl32, "glRenderMode")
	glRotated = doGetProcAddress(libopengl32, "glRotated")
	glRotatef = doGetProcAddress(libopengl32, "glRotatef")
	glScaled = doGetProcAddress(libopengl32, "glScaled")
	glScalef = doGetProcAddress(libopengl32, "glScalef")
	glScissor = doGetProcAddress(libopengl32, "glScissor")
	glSelectBuffer = doGetProcAddress(libopengl32, "glSelectBuffer")
	glShadeModel = doGetProcAddress(libopengl32, "glShadeModel")
	glStencilFunc = doGetProcAddress(libopengl32, "glStencilFunc")
	glStencilMask = doGetProcAddress(libopengl32, "glStencilMask")
	glStencilOp = doGetProcAddress(libopengl32, "glStencilOp")
	glTexCoord1d = doGetProcAddress(libopengl32, "glTexCoord1d")
	glTexCoord1dv = doGetProcAddress(libopengl32, "glTexCoord1dv")
	glTexCoord1f = doGetProcAddress(libopengl32, "glTexCoord1f")
	glTexCoord1fv = doGetProcAddress(libopengl32, "glTexCoord1fv")
	glTexCoord1i = doGetProcAddress(libopengl32, "glTexCoord1i")
	glTexCoord1iv = doGetProcAddress(libopengl32, "glTexCoord1iv")
	glTexCoord1s = doGetProcAddress(libopengl32, "glTexCoord1s")
	glTexCoord1sv = doGetProcAddress(libopengl32, "glTexCoord1sv")
	glTexCoord2d = doGetProcAddress(libopengl32, "glTexCoord2d")
	glTexCoord2dv = doGetProcAddress(libopengl32, "glTexCoord2dv")
	glTexCoord2f = doGetProcAddress(libopengl32, "glTexCoord2f")
	glTexCoord2fv = doGetProcAddress(libopengl32, "glTexCoord2fv")
	glTexCoord2i = doGetProcAddress(libopengl32, "glTexCoord2i")
	glTexCoord2iv = doGetProcAddress(libopengl32, "glTexCoord2iv")
	glTexCoord2s = doGetProcAddress(libopengl32, "glTexCoord2s")
	glTexCoord2sv = doGetProcAddress(libopengl32, "glTexCoord2sv")
	glTexCoord3d = doGetProcAddress(libopengl32, "glTexCoord3d")
	glTexCoord3dv = doGetProcAddress(libopengl32, "glTexCoord3dv")
	glTexCoord3f = doGetProcAddress(libopengl32, "glTexCoord3f")
	glTexCoord3fv = doGetProcAddress(libopengl32, "glTexCoord3fv")
	glTexCoord3i = doGetProcAddress(libopengl32, "glTexCoord3i")
	glTexCoord3iv = doGetProcAddress(libopengl32, "glTexCoord3iv")
	glTexCoord3s = doGetProcAddress(libopengl32, "glTexCoord3s")
	glTexCoord3sv = doGetProcAddress(libopengl32, "glTexCoord3sv")
	glTexCoord4d = doGetProcAddress(libopengl32, "glTexCoord4d")
	glTexCoord4dv = doGetProcAddress(libopengl32, "glTexCoord4dv")
	glTexCoord4f = doGetProcAddress(libopengl32, "glTexCoord4f")
	glTexCoord4fv = doGetProcAddress(libopengl32, "glTexCoord4fv")
	glTexCoord4i = doGetProcAddress(libopengl32, "glTexCoord4i")
	glTexCoord4iv = doGetProcAddress(libopengl32, "glTexCoord4iv")
	glTexCoord4s = doGetProcAddress(libopengl32, "glTexCoord4s")
	glTexCoord4sv = doGetProcAddress(libopengl32, "glTexCoord4sv")
	glTexCoordPointer = doGetProcAddress(libopengl32, "glTexCoordPointer")
	glTexEnvf = doGetProcAddress(libopengl32, "glTexEnvf")
	glTexEnvfv = doGetProcAddress(libopengl32, "glTexEnvfv")
	glTexEnvi = doGetProcAddress(libopengl32, "glTexEnvi")
	glTexEnviv = doGetProcAddress(libopengl32, "glTexEnviv")
	glTexGend = doGetProcAddress(libopengl32, "glTexGend")
	glTexGendv = doGetProcAddress(libopengl32, "glTexGendv")
	glTexGenf = doGetProcAddress(libopengl32, "glTexGenf")
	glTexGenfv = doGetProcAddress(libopengl32, "glTexGenfv")
	glTexGeni = doGetProcAddress(libopengl32, "glTexGeni")
	glTexGeniv = doGetProcAddress(libopengl32, "glTexGeniv")
	glTexImage1D = doGetProcAddress(libopengl32, "glTexImage1D")
	glTexImage2D = doGetProcAddress(libopengl32, "glTexImage2D")
	glTexParameterf = doGetProcAddress(libopengl32, "glTexParameterf")
	glTexParameterfv = doGetProcAddress(libopengl32, "glTexParameterfv")
	glTexParameteri = doGetProcAddress(libopengl32, "glTexParameteri")
	glTexParameteriv = doGetProcAddress(libopengl32, "glTexParameteriv")
	glTexSubImage1D = doGetProcAddress(libopengl32, "glTexSubImage1D")
	glTexSubImage2D = doGetProcAddress(libopengl32, "glTexSubImage2D")
	glTranslated = doGetProcAddress(libopengl32, "glTranslated")
	glTranslatef = doGetProcAddress(libopengl32, "glTranslatef")
	glVertex2d = doGetProcAddress(libopengl32, "glVertex2d")
	glVertex2dv = doGetProcAddress(libopengl32, "glVertex2dv")
	glVertex2f = doGetProcAddress(libopengl32, "glVertex2f")
	glVertex2fv = doGetProcAddress(libopengl32, "glVertex2fv")
	glVertex2i = doGetProcAddress(libopengl32, "glVertex2i")
	glVertex2iv = doGetProcAddress(libopengl32, "glVertex2iv")
	glVertex2s = doGetProcAddress(libopengl32, "glVertex2s")
	glVertex2sv = doGetProcAddress(libopengl32, "glVertex2sv")
	glVertex3d = doGetProcAddress(libopengl32, "glVertex3d")
	glVertex3dv = doGetProcAddress(libopengl32, "glVertex3dv")
	glVertex3f = doGetProcAddress(libopengl32, "glVertex3f")
	glVertex3fv = doGetProcAddress(libopengl32, "glVertex3fv")
	glVertex3i = doGetProcAddress(libopengl32, "glVertex3i")
	glVertex3iv = doGetProcAddress(libopengl32, "glVertex3iv")
	glVertex3s = doGetProcAddress(libopengl32, "glVertex3s")
	glVertex3sv = doGetProcAddress(libopengl32, "glVertex3sv")
	glVertex4d = doGetProcAddress(libopengl32, "glVertex4d")
	glVertex4dv = doGetProcAddress(libopengl32, "glVertex4dv")
	glVertex4f = doGetProcAddress(libopengl32, "glVertex4f")
	glVertex4fv = doGetProcAddress(libopengl32, "glVertex4fv")
	glVertex4i = doGetProcAddress(libopengl32, "glVertex4i")
	glVertex4iv = doGetProcAddress(libopengl32, "glVertex4iv")
	glVertex4s = doGetProcAddress(libopengl32, "glVertex4s")
	glVertex4sv = doGetProcAddress(libopengl32, "glVertex4sv")
	glVertexPointer = doGetProcAddress(libopengl32, "glVertexPointer")
	glViewport = doGetProcAddress(libopengl32, "glViewport")
	wglCopyContext = doGetProcAddress(libopengl32, "wglCopyContext")
	wglCreateContext = doGetProcAddress(libopengl32, "wglCreateContext")
	wglCreateLayerContext = doGetProcAddress(libopengl32, "wglCreateLayerContext")
	wglDeleteContext = doGetProcAddress(libopengl32, "wglDeleteContext")
	wglDescribeLayerPlane = doGetProcAddress(libopengl32, "wglDescribeLayerPlane")
	wglGetCurrentContext = doGetProcAddress(libopengl32, "wglGetCurrentContext")
	wglGetCurrentDC = doGetProcAddress(libopengl32, "wglGetCurrentDC")
	wglGetLayerPaletteEntries = doGetProcAddress(libopengl32, "wglGetLayerPaletteEntries")
	wglGetProcAddress = doGetProcAddress(libopengl32, "wglGetProcAddress")
	wglMakeCurrent = doGetProcAddress(libopengl32, "wglMakeCurrent")
	wglRealizeLayerPalette = doGetProcAddress(libopengl32, "wglRealizeLayerPalette")
	wglSetLayerPaletteEntries = doGetProcAddress(libopengl32, "wglSetLayerPaletteEntries")
	wglShareLists = doGetProcAddress(libopengl32, "wglShareLists")
	wglSwapLayerBuffers = doGetProcAddress(libopengl32, "wglSwapLayerBuffers")
	wglSwapMultipleBuffers = doGetProcAddress(libopengl32, "wglSwapMultipleBuffers")
	wglUseFontBitmaps = doGetProcAddress(libopengl32, "wglUseFontBitmapsW")
	wglUseFontOutlines = doGetProcAddress(libopengl32, "wglUseFontOutlinesW")
	glDebugEntry = doGetProcAddress(libopengl32, "glDebugEntry")
	wglChoosePixelFormat = doGetProcAddress(libopengl32, "wglChoosePixelFormat")
	wglDescribePixelFormat = doGetProcAddress(libopengl32, "wglDescribePixelFormat")
	wglGetPixelFormat = doGetProcAddress(libopengl32, "wglGetPixelFormat")
	wglSetPixelFormat = doGetProcAddress(libopengl32, "wglSetPixelFormat")
}

func GlAccum(op GLenum, value GLfloat) {
	syscall3(glAccum, 2,
		uintptr(op),
		uintptr(value),
		0)
}

func GlAlphaFunc(aFunc GLenum, ref GLclampf) {
	syscall3(glAlphaFunc, 2,
		uintptr(aFunc),
		uintptr(ref),
		0)
}

func GlAreTexturesResident(n GLsizei, textures /*const*/ *GLuint, residences *GLboolean) GLboolean {
	ret1 := syscall3(glAreTexturesResident, 3,
		uintptr(n),
		uintptr(unsafe.Pointer(textures)),
		uintptr(unsafe.Pointer(residences)))
	return GLboolean(ret1)
}

func GlArrayElement(i GLint) {
	syscall3(glArrayElement, 1,
		uintptr(i),
		0,
		0)
}

func GlBegin(mode GLenum) {
	syscall3(glBegin, 1,
		uintptr(mode),
		0,
		0)
}

func GlBindTexture(target GLenum, texture GLuint) {
	syscall3(glBindTexture, 2,
		uintptr(target),
		uintptr(texture),
		0)
}

func GlBitmap(width GLsizei, height GLsizei, xorig GLfloat, yorig GLfloat, xmove GLfloat, ymove GLfloat, bitmap /*const*/ *GLubyte) {
	syscall9(glBitmap, 7,
		uintptr(width),
		uintptr(height),
		uintptr(xorig),
		uintptr(yorig),
		uintptr(xmove),
		uintptr(ymove),
		uintptr(unsafe.Pointer(bitmap)),
		0,
		0)
}

func GlBlendFunc(sfactor GLenum, dfactor GLenum) {
	syscall3(glBlendFunc, 2,
		uintptr(sfactor),
		uintptr(dfactor),
		0)
}

func GlCallList(list GLuint) {
	syscall3(glCallList, 1,
		uintptr(list),
		0,
		0)
}

func GlCallLists(n GLsizei, aType GLenum, lists /*const*/ uintptr) {
	syscall3(glCallLists, 3,
		uintptr(n),
		uintptr(aType),
		lists)
}

func GlClear(mask GLbitfield) {
	syscall3(glClear, 1,
		uintptr(mask),
		0,
		0)
}

func GlClearAccum(red GLfloat, green GLfloat, blue GLfloat, alpha GLfloat) {
	syscall6(glClearAccum, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	syscall6(glClearColor, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlClearDepth(depth GLclampd) {
	syscall3(glClearDepth, 1,
		uintptr(depth),
		0,
		0)
}

func GlClearIndex(c GLfloat) {
	syscall3(glClearIndex, 1,
		uintptr(c),
		0,
		0)
}

func GlClearStencil(s GLint) {
	syscall3(glClearStencil, 1,
		uintptr(s),
		0,
		0)
}

func GlClipPlane(plane GLenum, equation /*const*/ *GLdouble) {
	syscall3(glClipPlane, 2,
		uintptr(plane),
		uintptr(unsafe.Pointer(equation)),
		0)
}

func GlColor3b(red GLbyte, green GLbyte, blue GLbyte) {
	syscall3(glColor3b, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3bv(v /*const*/ *GLbyte) {
	syscall3(glColor3bv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3d(red GLdouble, green GLdouble, blue GLdouble) {
	syscall3(glColor3d, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3dv(v /*const*/ *GLdouble) {
	syscall3(glColor3dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3f(red GLfloat, green GLfloat, blue GLfloat) {
	syscall3(glColor3f, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3fv(v /*const*/ *GLfloat) {
	syscall3(glColor3fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3i(red GLint, green GLint, blue GLint) {
	syscall3(glColor3i, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3iv(v /*const*/ *GLint) {
	syscall3(glColor3iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3s(red GLshort, green GLshort, blue GLshort) {
	syscall3(glColor3s, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3sv(v /*const*/ *GLshort) {
	syscall3(glColor3sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3ub(red GLubyte, green GLubyte, blue GLubyte) {
	syscall3(glColor3ub, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3ubv(v /*const*/ *GLubyte) {
	syscall3(glColor3ubv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3ui(red GLuint, green GLuint, blue GLuint) {
	syscall3(glColor3ui, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3uiv(v /*const*/ *GLuint) {
	syscall3(glColor3uiv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor3us(red GLushort, green GLushort, blue GLushort) {
	syscall3(glColor3us, 3,
		uintptr(red),
		uintptr(green),
		uintptr(blue))
}

func GlColor3usv(v /*const*/ *GLushort) {
	syscall3(glColor3usv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4b(red GLbyte, green GLbyte, blue GLbyte, alpha GLbyte) {
	syscall6(glColor4b, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4bv(v /*const*/ *GLbyte) {
	syscall3(glColor4bv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4d(red GLdouble, green GLdouble, blue GLdouble, alpha GLdouble) {
	syscall6(glColor4d, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4dv(v /*const*/ *GLdouble) {
	syscall3(glColor4dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4f(red GLfloat, green GLfloat, blue GLfloat, alpha GLfloat) {
	syscall6(glColor4f, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4fv(v /*const*/ *GLfloat) {
	syscall3(glColor4fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4i(red GLint, green GLint, blue GLint, alpha GLint) {
	syscall6(glColor4i, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4iv(v /*const*/ *GLint) {
	syscall3(glColor4iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4s(red GLshort, green GLshort, blue GLshort, alpha GLshort) {
	syscall6(glColor4s, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4sv(v /*const*/ *GLshort) {
	syscall3(glColor4sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4ub(red GLubyte, green GLubyte, blue GLubyte, alpha GLubyte) {
	syscall6(glColor4ub, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4ubv(v /*const*/ *GLubyte) {
	syscall3(glColor4ubv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4ui(red GLuint, green GLuint, blue GLuint, alpha GLuint) {
	syscall6(glColor4ui, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4uiv(v /*const*/ *GLuint) {
	syscall3(glColor4uiv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColor4us(red GLushort, green GLushort, blue GLushort, alpha GLushort) {
	syscall6(glColor4us, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColor4usv(v /*const*/ *GLushort) {
	syscall3(glColor4usv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlColorMask(red GLboolean, green GLboolean, blue GLboolean, alpha GLboolean) {
	syscall6(glColorMask, 4,
		uintptr(red),
		uintptr(green),
		uintptr(blue),
		uintptr(alpha),
		0,
		0)
}

func GlColorMaterial(face GLenum, mode GLenum) {
	syscall3(glColorMaterial, 2,
		uintptr(face),
		uintptr(mode),
		0)
}

func GlColorPointer(size GLint, aType GLenum, stride GLsizei, pointer /*const*/ uintptr) {
	syscall6(glColorPointer, 4,
		uintptr(size),
		uintptr(aType),
		uintptr(stride),
		pointer,
		0,
		0)
}

func GlCopyPixels(x GLint, y GLint, width GLsizei, height GLsizei, aType GLenum) {
	syscall6(glCopyPixels, 5,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(aType),
		0)
}

func GlCopyTexImage1D(target GLenum, level GLint, internalFormat GLenum, x GLint, y GLint, width GLsizei, border GLint) {
	syscall9(glCopyTexImage1D, 7,
		uintptr(target),
		uintptr(level),
		uintptr(internalFormat),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(border),
		0,
		0)
}

func GlCopyTexImage2D(target GLenum, level GLint, internalFormat GLenum, x GLint, y GLint, width GLsizei, height GLsizei, border GLint) {
	syscall9(glCopyTexImage2D, 8,
		uintptr(target),
		uintptr(level),
		uintptr(internalFormat),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(border),
		0)
}

func GlCopyTexSubImage1D(target GLenum, level GLint, xoffset GLint, x GLint, y GLint, width GLsizei) {
	syscall6(glCopyTexSubImage1D, 6,
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(x),
		uintptr(y),
		uintptr(width))
}

func GlCopyTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, x GLint, y GLint, width GLsizei, height GLsizei) {
	syscall9(glCopyTexSubImage2D, 8,
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(yoffset),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		0)
}

func GlCullFace(mode GLenum) {
	syscall3(glCullFace, 1,
		uintptr(mode),
		0,
		0)
}

func GlDeleteLists(list GLuint, aRange GLsizei) {
	syscall3(glDeleteLists, 2,
		uintptr(list),
		uintptr(aRange),
		0)
}

func GlDeleteTextures(n GLsizei, textures /*const*/ *GLuint) {
	syscall3(glDeleteTextures, 2,
		uintptr(n),
		uintptr(unsafe.Pointer(textures)),
		0)
}

func GlDepthFunc(aFunc GLenum) {
	syscall3(glDepthFunc, 1,
		uintptr(aFunc),
		0,
		0)
}

func GlDepthMask(flag GLboolean) {
	syscall3(glDepthMask, 1,
		uintptr(flag),
		0,
		0)
}

func GlDepthRange(zNear GLclampd, zFar GLclampd) {
	syscall3(glDepthRange, 2,
		uintptr(zNear),
		uintptr(zFar),
		0)
}

func GlDisable(cap GLenum) {
	syscall3(glDisable, 1,
		uintptr(cap),
		0,
		0)
}

func GlDisableClientState(array GLenum) {
	syscall3(glDisableClientState, 1,
		uintptr(array),
		0,
		0)
}

func GlDrawArrays(mode GLenum, first GLint, count GLsizei) {
	syscall3(glDrawArrays, 3,
		uintptr(mode),
		uintptr(first),
		uintptr(count))
}

func GlDrawBuffer(mode GLenum) {
	syscall3(glDrawBuffer, 1,
		uintptr(mode),
		0,
		0)
}

func GlDrawElements(mode GLenum, count GLsizei, aType GLenum, indices /*const*/ uintptr) {
	syscall6(glDrawElements, 4,
		uintptr(mode),
		uintptr(count),
		uintptr(aType),
		indices,
		0,
		0)
}

func GlDrawPixels(width GLsizei, height GLsizei, format GLenum, aType GLenum, pixels /*const*/ uintptr) {
	syscall6(glDrawPixels, 5,
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(aType),
		pixels,
		0)
}

func GlEdgeFlag(flag GLboolean) {
	syscall3(glEdgeFlag, 1,
		uintptr(flag),
		0,
		0)
}

func GlEdgeFlagPointer(stride GLsizei, pointer /*const*/ uintptr) {
	syscall3(glEdgeFlagPointer, 2,
		uintptr(stride),
		pointer,
		0)
}

func GlEdgeFlagv(flag /*const*/ *GLboolean) {
	syscall3(glEdgeFlagv, 1,
		uintptr(unsafe.Pointer(flag)),
		0,
		0)
}

func GlEnable(cap GLenum) {
	syscall3(glEnable, 1,
		uintptr(cap),
		0,
		0)
}

func GlEnableClientState(array GLenum) {
	syscall3(glEnableClientState, 1,
		uintptr(array),
		0,
		0)
}

func GlEnd() {
	syscall3(glEnd, 0,
		0,
		0,
		0)
}

func GlEndList() {
	syscall3(glEndList, 0,
		0,
		0,
		0)
}

func GlEvalCoord1d(u GLdouble) {
	syscall3(glEvalCoord1d, 1,
		uintptr(u),
		0,
		0)
}

func GlEvalCoord1dv(u /*const*/ *GLdouble) {
	syscall3(glEvalCoord1dv, 1,
		uintptr(unsafe.Pointer(u)),
		0,
		0)
}

func GlEvalCoord1f(u GLfloat) {
	syscall3(glEvalCoord1f, 1,
		uintptr(u),
		0,
		0)
}

func GlEvalCoord1fv(u /*const*/ *GLfloat) {
	syscall3(glEvalCoord1fv, 1,
		uintptr(unsafe.Pointer(u)),
		0,
		0)
}

func GlEvalCoord2d(u GLdouble, v GLdouble) {
	syscall3(glEvalCoord2d, 2,
		uintptr(u),
		uintptr(v),
		0)
}

func GlEvalCoord2dv(u /*const*/ *GLdouble) {
	syscall3(glEvalCoord2dv, 1,
		uintptr(unsafe.Pointer(u)),
		0,
		0)
}

func GlEvalCoord2f(u GLfloat, v GLfloat) {
	syscall3(glEvalCoord2f, 2,
		uintptr(u),
		uintptr(v),
		0)
}

func GlEvalCoord2fv(u /*const*/ *GLfloat) {
	syscall3(glEvalCoord2fv, 1,
		uintptr(unsafe.Pointer(u)),
		0,
		0)
}

func GlEvalMesh1(mode GLenum, i1 GLint, i2 GLint) {
	syscall3(glEvalMesh1, 3,
		uintptr(mode),
		uintptr(i1),
		uintptr(i2))
}

func GlEvalMesh2(mode GLenum, i1 GLint, i2 GLint, j1 GLint, j2 GLint) {
	syscall6(glEvalMesh2, 5,
		uintptr(mode),
		uintptr(i1),
		uintptr(i2),
		uintptr(j1),
		uintptr(j2),
		0)
}

func GlEvalPoint1(i GLint) {
	syscall3(glEvalPoint1, 1,
		uintptr(i),
		0,
		0)
}

func GlEvalPoint2(i GLint, j GLint) {
	syscall3(glEvalPoint2, 2,
		uintptr(i),
		uintptr(j),
		0)
}

func GlFeedbackBuffer(size GLsizei, aType GLenum, buffer *GLfloat) {
	syscall3(glFeedbackBuffer, 3,
		uintptr(size),
		uintptr(aType),
		uintptr(unsafe.Pointer(buffer)))
}

func GlFinish() {
	syscall3(glFinish, 0,
		0,
		0,
		0)
}

func GlFlush() {
	syscall3(glFlush, 0,
		0,
		0,
		0)
}

func GlFogf(pname GLenum, param GLfloat) {
	syscall3(glFogf, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlFogfv(pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glFogfv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlFogi(pname GLenum, param GLint) {
	syscall3(glFogi, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlFogiv(pname GLenum, params /*const*/ *GLint) {
	syscall3(glFogiv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlFrontFace(mode GLenum) {
	syscall3(glFrontFace, 1,
		uintptr(mode),
		0,
		0)
}

func GlFrustum(left GLdouble, right GLdouble, bottom GLdouble, top GLdouble, zNear GLdouble, zFar GLdouble) {
	syscall6(glFrustum, 6,
		uintptr(left),
		uintptr(right),
		uintptr(bottom),
		uintptr(top),
		uintptr(zNear),
		uintptr(zFar))
}

func GlGenLists(aRange GLsizei) GLuint {
	ret1 := syscall3(glGenLists, 1,
		uintptr(aRange),
		0,
		0)
	return GLuint(ret1)
}

func GlGenTextures(n GLsizei, textures *GLuint) {
	syscall3(glGenTextures, 2,
		uintptr(n),
		uintptr(unsafe.Pointer(textures)),
		0)
}

func GlGetBooleanv(pname GLenum, params *GLboolean) {
	syscall3(glGetBooleanv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlGetClipPlane(plane GLenum, equation *GLdouble) {
	syscall3(glGetClipPlane, 2,
		uintptr(plane),
		uintptr(unsafe.Pointer(equation)),
		0)
}

func GlGetDoublev(pname GLenum, params *GLdouble) {
	syscall3(glGetDoublev, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlGetError() GLenum {
	ret1 := syscall3(glGetError, 0,
		0,
		0,
		0)
	return GLenum(ret1)
}

func GlGetFloatv(pname GLenum, params *GLfloat) {
	syscall3(glGetFloatv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlGetIntegerv(pname GLenum, params *GLint) {
	syscall3(glGetIntegerv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlGetLightfv(light GLenum, pname GLenum, params *GLfloat) {
	syscall3(glGetLightfv, 3,
		uintptr(light),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetLightiv(light GLenum, pname GLenum, params *GLint) {
	syscall3(glGetLightiv, 3,
		uintptr(light),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetMapdv(target GLenum, query GLenum, v *GLdouble) {
	syscall3(glGetMapdv, 3,
		uintptr(target),
		uintptr(query),
		uintptr(unsafe.Pointer(v)))
}

func GlGetMapfv(target GLenum, query GLenum, v *GLfloat) {
	syscall3(glGetMapfv, 3,
		uintptr(target),
		uintptr(query),
		uintptr(unsafe.Pointer(v)))
}

func GlGetMapiv(target GLenum, query GLenum, v *GLint) {
	syscall3(glGetMapiv, 3,
		uintptr(target),
		uintptr(query),
		uintptr(unsafe.Pointer(v)))
}

func GlGetMaterialfv(face GLenum, pname GLenum, params *GLfloat) {
	syscall3(glGetMaterialfv, 3,
		uintptr(face),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetMaterialiv(face GLenum, pname GLenum, params *GLint) {
	syscall3(glGetMaterialiv, 3,
		uintptr(face),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetPixelMapfv(aMap GLenum, values *GLfloat) {
	syscall3(glGetPixelMapfv, 2,
		uintptr(aMap),
		uintptr(unsafe.Pointer(values)),
		0)
}

func GlGetPixelMapuiv(aMap GLenum, values *GLuint) {
	syscall3(glGetPixelMapuiv, 2,
		uintptr(aMap),
		uintptr(unsafe.Pointer(values)),
		0)
}

func GlGetPixelMapusv(aMap GLenum, values *GLushort) {
	syscall3(glGetPixelMapusv, 2,
		uintptr(aMap),
		uintptr(unsafe.Pointer(values)),
		0)
}

func GlGetPointerv(pname GLenum, params uintptr) {
	syscall3(glGetPointerv, 2,
		uintptr(pname),
		params,
		0)
}

func GlGetPolygonStipple(mask *GLubyte) {
	syscall3(glGetPolygonStipple, 1,
		uintptr(unsafe.Pointer(mask)),
		0,
		0)
}

func GlGetTexEnvfv(target GLenum, pname GLenum, params *GLfloat) {
	syscall3(glGetTexEnvfv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetTexEnviv(target GLenum, pname GLenum, params *GLint) {
	syscall3(glGetTexEnviv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetTexGendv(coord GLenum, pname GLenum, params *GLdouble) {
	syscall3(glGetTexGendv, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetTexGenfv(coord GLenum, pname GLenum, params *GLfloat) {
	syscall3(glGetTexGenfv, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetTexGeniv(coord GLenum, pname GLenum, params *GLint) {
	syscall3(glGetTexGeniv, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetTexImage(target GLenum, level GLint, format GLenum, aType GLenum, pixels uintptr) {
	syscall6(glGetTexImage, 5,
		uintptr(target),
		uintptr(level),
		uintptr(format),
		uintptr(aType),
		pixels,
		0)
}

func GlGetTexLevelParameterfv(target GLenum, level GLint, pname GLenum, params *GLfloat) {
	syscall6(glGetTexLevelParameterfv, 4,
		uintptr(target),
		uintptr(level),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0,
		0)
}

func GlGetTexLevelParameteriv(target GLenum, level GLint, pname GLenum, params *GLint) {
	syscall6(glGetTexLevelParameteriv, 4,
		uintptr(target),
		uintptr(level),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0,
		0)
}

func GlGetTexParameterfv(target GLenum, pname GLenum, params *GLfloat) {
	syscall3(glGetTexParameterfv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlGetTexParameteriv(target GLenum, pname GLenum, params *GLint) {
	syscall3(glGetTexParameteriv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlHint(target GLenum, mode GLenum) {
	syscall3(glHint, 2,
		uintptr(target),
		uintptr(mode),
		0)
}

func GlIndexMask(mask GLuint) {
	syscall3(glIndexMask, 1,
		uintptr(mask),
		0,
		0)
}

func GlIndexPointer(aType GLenum, stride GLsizei, pointer /*const*/ uintptr) {
	syscall3(glIndexPointer, 3,
		uintptr(aType),
		uintptr(stride),
		pointer)
}

func GlIndexd(c GLdouble) {
	syscall3(glIndexd, 1,
		uintptr(c),
		0,
		0)
}

func GlIndexdv(c /*const*/ *GLdouble) {
	syscall3(glIndexdv, 1,
		uintptr(unsafe.Pointer(c)),
		0,
		0)
}

func GlIndexf(c GLfloat) {
	syscall3(glIndexf, 1,
		uintptr(c),
		0,
		0)
}

func GlIndexfv(c /*const*/ *GLfloat) {
	syscall3(glIndexfv, 1,
		uintptr(unsafe.Pointer(c)),
		0,
		0)
}

func GlIndexi(c GLint) {
	syscall3(glIndexi, 1,
		uintptr(c),
		0,
		0)
}

func GlIndexiv(c /*const*/ *GLint) {
	syscall3(glIndexiv, 1,
		uintptr(unsafe.Pointer(c)),
		0,
		0)
}

func GlIndexs(c GLshort) {
	syscall3(glIndexs, 1,
		uintptr(c),
		0,
		0)
}

func GlIndexsv(c /*const*/ *GLshort) {
	syscall3(glIndexsv, 1,
		uintptr(unsafe.Pointer(c)),
		0,
		0)
}

func GlIndexub(c GLubyte) {
	syscall3(glIndexub, 1,
		uintptr(c),
		0,
		0)
}

func GlIndexubv(c /*const*/ *GLubyte) {
	syscall3(glIndexubv, 1,
		uintptr(unsafe.Pointer(c)),
		0,
		0)
}

func GlInitNames() {
	syscall3(glInitNames, 0,
		0,
		0,
		0)
}

func GlInterleavedArrays(format GLenum, stride GLsizei, pointer /*const*/ uintptr) {
	syscall3(glInterleavedArrays, 3,
		uintptr(format),
		uintptr(stride),
		pointer)
}

func GlIsEnabled(cap GLenum) GLboolean {
	ret1 := syscall3(glIsEnabled, 1,
		uintptr(cap),
		0,
		0)
	return GLboolean(ret1)
}

func GlIsList(list GLuint) GLboolean {
	ret1 := syscall3(glIsList, 1,
		uintptr(list),
		0,
		0)
	return GLboolean(ret1)
}

func GlIsTexture(texture GLuint) GLboolean {
	ret1 := syscall3(glIsTexture, 1,
		uintptr(texture),
		0,
		0)
	return GLboolean(ret1)
}

func GlLightModelf(pname GLenum, param GLfloat) {
	syscall3(glLightModelf, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlLightModelfv(pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glLightModelfv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlLightModeli(pname GLenum, param GLint) {
	syscall3(glLightModeli, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlLightModeliv(pname GLenum, params /*const*/ *GLint) {
	syscall3(glLightModeliv, 2,
		uintptr(pname),
		uintptr(unsafe.Pointer(params)),
		0)
}

func GlLightf(light GLenum, pname GLenum, param GLfloat) {
	syscall3(glLightf, 3,
		uintptr(light),
		uintptr(pname),
		uintptr(param))
}

func GlLightfv(light GLenum, pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glLightfv, 3,
		uintptr(light),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlLighti(light GLenum, pname GLenum, param GLint) {
	syscall3(glLighti, 3,
		uintptr(light),
		uintptr(pname),
		uintptr(param))
}

func GlLightiv(light GLenum, pname GLenum, params /*const*/ *GLint) {
	syscall3(glLightiv, 3,
		uintptr(light),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlLineStipple(factor GLint, pattern GLushort) {
	syscall3(glLineStipple, 2,
		uintptr(factor),
		uintptr(pattern),
		0)
}

func GlLineWidth(width GLfloat) {
	syscall3(glLineWidth, 1,
		uintptr(width),
		0,
		0)
}

func GlListBase(base GLuint) {
	syscall3(glListBase, 1,
		uintptr(base),
		0,
		0)
}

func GlLoadIdentity() {
	syscall3(glLoadIdentity, 0,
		0,
		0,
		0)
}

func GlLoadMatrixd(m /*const*/ *GLdouble) {
	syscall3(glLoadMatrixd, 1,
		uintptr(unsafe.Pointer(m)),
		0,
		0)
}

func GlLoadMatrixf(m /*const*/ *GLfloat) {
	syscall3(glLoadMatrixf, 1,
		uintptr(unsafe.Pointer(m)),
		0,
		0)
}

func GlLoadName(name GLuint) {
	syscall3(glLoadName, 1,
		uintptr(name),
		0,
		0)
}

func GlLogicOp(opcode GLenum) {
	syscall3(glLogicOp, 1,
		uintptr(opcode),
		0,
		0)
}

func GlMap1d(target GLenum, u1 GLdouble, u2 GLdouble, stride GLint, order GLint, points /*const*/ *GLdouble) {
	syscall6(glMap1d, 6,
		uintptr(target),
		uintptr(u1),
		uintptr(u2),
		uintptr(stride),
		uintptr(order),
		uintptr(unsafe.Pointer(points)))
}

func GlMap1f(target GLenum, u1 GLfloat, u2 GLfloat, stride GLint, order GLint, points /*const*/ *GLfloat) {
	syscall6(glMap1f, 6,
		uintptr(target),
		uintptr(u1),
		uintptr(u2),
		uintptr(stride),
		uintptr(order),
		uintptr(unsafe.Pointer(points)))
}

func GlMap2d(target GLenum, u1 GLdouble, u2 GLdouble, ustride GLint, uorder GLint, v1 GLdouble, v2 GLdouble, vstride GLint, vorder GLint, points /*const*/ *GLdouble) {
	syscall12(glMap2d, 10,
		uintptr(target),
		uintptr(u1),
		uintptr(u2),
		uintptr(ustride),
		uintptr(uorder),
		uintptr(v1),
		uintptr(v2),
		uintptr(vstride),
		uintptr(vorder),
		uintptr(unsafe.Pointer(points)),
		0,
		0)
}

func GlMap2f(target GLenum, u1 GLfloat, u2 GLfloat, ustride GLint, uorder GLint, v1 GLfloat, v2 GLfloat, vstride GLint, vorder GLint, points /*const*/ *GLfloat) {
	syscall12(glMap2f, 10,
		uintptr(target),
		uintptr(u1),
		uintptr(u2),
		uintptr(ustride),
		uintptr(uorder),
		uintptr(v1),
		uintptr(v2),
		uintptr(vstride),
		uintptr(vorder),
		uintptr(unsafe.Pointer(points)),
		0,
		0)
}

func GlMapGrid1d(un GLint, u1 GLdouble, u2 GLdouble) {
	syscall3(glMapGrid1d, 3,
		uintptr(un),
		uintptr(u1),
		uintptr(u2))
}

func GlMapGrid1f(un GLint, u1 GLfloat, u2 GLfloat) {
	syscall3(glMapGrid1f, 3,
		uintptr(un),
		uintptr(u1),
		uintptr(u2))
}

func GlMapGrid2d(un GLint, u1 GLdouble, u2 GLdouble, vn GLint, v1 GLdouble, v2 GLdouble) {
	syscall6(glMapGrid2d, 6,
		uintptr(un),
		uintptr(u1),
		uintptr(u2),
		uintptr(vn),
		uintptr(v1),
		uintptr(v2))
}

func GlMapGrid2f(un GLint, u1 GLfloat, u2 GLfloat, vn GLint, v1 GLfloat, v2 GLfloat) {
	syscall6(glMapGrid2f, 6,
		uintptr(un),
		uintptr(u1),
		uintptr(u2),
		uintptr(vn),
		uintptr(v1),
		uintptr(v2))
}

func GlMaterialf(face GLenum, pname GLenum, param GLfloat) {
	syscall3(glMaterialf, 3,
		uintptr(face),
		uintptr(pname),
		uintptr(param))
}

func GlMaterialfv(face GLenum, pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glMaterialfv, 3,
		uintptr(face),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlMateriali(face GLenum, pname GLenum, param GLint) {
	syscall3(glMateriali, 3,
		uintptr(face),
		uintptr(pname),
		uintptr(param))
}

func GlMaterialiv(face GLenum, pname GLenum, params /*const*/ *GLint) {
	syscall3(glMaterialiv, 3,
		uintptr(face),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlMatrixMode(mode GLenum) {
	syscall3(glMatrixMode, 1,
		uintptr(mode),
		0,
		0)
}

func GlMultMatrixd(m /*const*/ *GLdouble) {
	syscall3(glMultMatrixd, 1,
		uintptr(unsafe.Pointer(m)),
		0,
		0)
}

func GlMultMatrixf(m /*const*/ *GLfloat) {
	syscall3(glMultMatrixf, 1,
		uintptr(unsafe.Pointer(m)),
		0,
		0)
}

func GlNewList(list GLuint, mode GLenum) {
	syscall3(glNewList, 2,
		uintptr(list),
		uintptr(mode),
		0)
}

func GlNormal3b(nx GLbyte, ny GLbyte, nz GLbyte) {
	syscall3(glNormal3b, 3,
		uintptr(nx),
		uintptr(ny),
		uintptr(nz))
}

func GlNormal3bv(v /*const*/ *GLbyte) {
	syscall3(glNormal3bv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlNormal3d(nx GLdouble, ny GLdouble, nz GLdouble) {
	syscall3(glNormal3d, 3,
		uintptr(nx),
		uintptr(ny),
		uintptr(nz))
}

func GlNormal3dv(v /*const*/ *GLdouble) {
	syscall3(glNormal3dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlNormal3f(nx GLfloat, ny GLfloat, nz GLfloat) {
	syscall3(glNormal3f, 3,
		uintptr(nx),
		uintptr(ny),
		uintptr(nz))
}

func GlNormal3fv(v /*const*/ *GLfloat) {
	syscall3(glNormal3fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlNormal3i(nx GLint, ny GLint, nz GLint) {
	syscall3(glNormal3i, 3,
		uintptr(nx),
		uintptr(ny),
		uintptr(nz))
}

func GlNormal3iv(v /*const*/ *GLint) {
	syscall3(glNormal3iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlNormal3s(nx GLshort, ny GLshort, nz GLshort) {
	syscall3(glNormal3s, 3,
		uintptr(nx),
		uintptr(ny),
		uintptr(nz))
}

func GlNormal3sv(v /*const*/ *GLshort) {
	syscall3(glNormal3sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlNormalPointer(aType GLenum, stride GLsizei, pointer /*const*/ uintptr) {
	syscall3(glNormalPointer, 3,
		uintptr(aType),
		uintptr(stride),
		pointer)
}

func GlOrtho(left GLdouble, right GLdouble, bottom GLdouble, top GLdouble, zNear GLdouble, zFar GLdouble) {
	syscall6(glOrtho, 6,
		uintptr(left),
		uintptr(right),
		uintptr(bottom),
		uintptr(top),
		uintptr(zNear),
		uintptr(zFar))
}

func GlPassThrough(token GLfloat) {
	syscall3(glPassThrough, 1,
		uintptr(token),
		0,
		0)
}

func GlPixelMapfv(aMap GLenum, mapsize GLsizei, values /*const*/ *GLfloat) {
	syscall3(glPixelMapfv, 3,
		uintptr(aMap),
		uintptr(mapsize),
		uintptr(unsafe.Pointer(values)))
}

func GlPixelMapuiv(aMap GLenum, mapsize GLsizei, values /*const*/ *GLuint) {
	syscall3(glPixelMapuiv, 3,
		uintptr(aMap),
		uintptr(mapsize),
		uintptr(unsafe.Pointer(values)))
}

func GlPixelMapusv(aMap GLenum, mapsize GLsizei, values /*const*/ *GLushort) {
	syscall3(glPixelMapusv, 3,
		uintptr(aMap),
		uintptr(mapsize),
		uintptr(unsafe.Pointer(values)))
}

func GlPixelStoref(pname GLenum, param GLfloat) {
	syscall3(glPixelStoref, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlPixelStorei(pname GLenum, param GLint) {
	syscall3(glPixelStorei, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlPixelTransferf(pname GLenum, param GLfloat) {
	syscall3(glPixelTransferf, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlPixelTransferi(pname GLenum, param GLint) {
	syscall3(glPixelTransferi, 2,
		uintptr(pname),
		uintptr(param),
		0)
}

func GlPixelZoom(xfactor GLfloat, yfactor GLfloat) {
	syscall3(glPixelZoom, 2,
		uintptr(xfactor),
		uintptr(yfactor),
		0)
}

func GlPointSize(size GLfloat) {
	syscall3(glPointSize, 1,
		uintptr(size),
		0,
		0)
}

func GlPolygonMode(face GLenum, mode GLenum) {
	syscall3(glPolygonMode, 2,
		uintptr(face),
		uintptr(mode),
		0)
}

func GlPolygonOffset(factor GLfloat, units GLfloat) {
	syscall3(glPolygonOffset, 2,
		uintptr(factor),
		uintptr(units),
		0)
}

func GlPolygonStipple(mask /*const*/ *GLubyte) {
	syscall3(glPolygonStipple, 1,
		uintptr(unsafe.Pointer(mask)),
		0,
		0)
}

func GlPopAttrib() {
	syscall3(glPopAttrib, 0,
		0,
		0,
		0)
}

func GlPopClientAttrib() {
	syscall3(glPopClientAttrib, 0,
		0,
		0,
		0)
}

func GlPopMatrix() {
	syscall3(glPopMatrix, 0,
		0,
		0,
		0)
}

func GlPopName() {
	syscall3(glPopName, 0,
		0,
		0,
		0)
}

func GlPrioritizeTextures(n GLsizei, textures /*const*/ *GLuint, priorities /*const*/ *GLclampf) {
	syscall3(glPrioritizeTextures, 3,
		uintptr(n),
		uintptr(unsafe.Pointer(textures)),
		uintptr(unsafe.Pointer(priorities)))
}

func GlPushAttrib(mask GLbitfield) {
	syscall3(glPushAttrib, 1,
		uintptr(mask),
		0,
		0)
}

func GlPushClientAttrib(mask GLbitfield) {
	syscall3(glPushClientAttrib, 1,
		uintptr(mask),
		0,
		0)
}

func GlPushMatrix() {
	syscall3(glPushMatrix, 0,
		0,
		0,
		0)
}

func GlPushName(name GLuint) {
	syscall3(glPushName, 1,
		uintptr(name),
		0,
		0)
}

func GlRasterPos2d(x GLdouble, y GLdouble) {
	syscall3(glRasterPos2d, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlRasterPos2dv(v /*const*/ *GLdouble) {
	syscall3(glRasterPos2dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos2f(x GLfloat, y GLfloat) {
	syscall3(glRasterPos2f, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlRasterPos2fv(v /*const*/ *GLfloat) {
	syscall3(glRasterPos2fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos2i(x GLint, y GLint) {
	syscall3(glRasterPos2i, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlRasterPos2iv(v /*const*/ *GLint) {
	syscall3(glRasterPos2iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos2s(x GLshort, y GLshort) {
	syscall3(glRasterPos2s, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlRasterPos2sv(v /*const*/ *GLshort) {
	syscall3(glRasterPos2sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos3d(x GLdouble, y GLdouble, z GLdouble) {
	syscall3(glRasterPos3d, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlRasterPos3dv(v /*const*/ *GLdouble) {
	syscall3(glRasterPos3dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos3f(x GLfloat, y GLfloat, z GLfloat) {
	syscall3(glRasterPos3f, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlRasterPos3fv(v /*const*/ *GLfloat) {
	syscall3(glRasterPos3fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos3i(x GLint, y GLint, z GLint) {
	syscall3(glRasterPos3i, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlRasterPos3iv(v /*const*/ *GLint) {
	syscall3(glRasterPos3iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos3s(x GLshort, y GLshort, z GLshort) {
	syscall3(glRasterPos3s, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlRasterPos3sv(v /*const*/ *GLshort) {
	syscall3(glRasterPos3sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos4d(x GLdouble, y GLdouble, z GLdouble, w GLdouble) {
	syscall6(glRasterPos4d, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlRasterPos4dv(v /*const*/ *GLdouble) {
	syscall3(glRasterPos4dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos4f(x GLfloat, y GLfloat, z GLfloat, w GLfloat) {
	syscall6(glRasterPos4f, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlRasterPos4fv(v /*const*/ *GLfloat) {
	syscall3(glRasterPos4fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos4i(x GLint, y GLint, z GLint, w GLint) {
	syscall6(glRasterPos4i, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlRasterPos4iv(v /*const*/ *GLint) {
	syscall3(glRasterPos4iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlRasterPos4s(x GLshort, y GLshort, z GLshort, w GLshort) {
	syscall6(glRasterPos4s, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlRasterPos4sv(v /*const*/ *GLshort) {
	syscall3(glRasterPos4sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlReadBuffer(mode GLenum) {
	syscall3(glReadBuffer, 1,
		uintptr(mode),
		0,
		0)
}

func GlReadPixels(x GLint, y GLint, width GLsizei, height GLsizei, format GLenum, aType GLenum, pixels uintptr) {
	syscall9(glReadPixels, 7,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(aType),
		pixels,
		0,
		0)
}

func GlRectd(x1 GLdouble, y1 GLdouble, x2 GLdouble, y2 GLdouble) {
	syscall6(glRectd, 4,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		0,
		0)
}

func GlRectdv(v1 /*const*/ *GLdouble, v2 /*const*/ *GLdouble) {
	syscall3(glRectdv, 2,
		uintptr(unsafe.Pointer(v1)),
		uintptr(unsafe.Pointer(v2)),
		0)
}

func GlRectf(x1 GLfloat, y1 GLfloat, x2 GLfloat, y2 GLfloat) {
	syscall6(glRectf, 4,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		0,
		0)
}

func GlRectfv(v1 /*const*/ *GLfloat, v2 /*const*/ *GLfloat) {
	syscall3(glRectfv, 2,
		uintptr(unsafe.Pointer(v1)),
		uintptr(unsafe.Pointer(v2)),
		0)
}

func GlRecti(x1 GLint, y1 GLint, x2 GLint, y2 GLint) {
	syscall6(glRecti, 4,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		0,
		0)
}

func GlRectiv(v1 /*const*/ *GLint, v2 /*const*/ *GLint) {
	syscall3(glRectiv, 2,
		uintptr(unsafe.Pointer(v1)),
		uintptr(unsafe.Pointer(v2)),
		0)
}

func GlRects(x1 GLshort, y1 GLshort, x2 GLshort, y2 GLshort) {
	syscall6(glRects, 4,
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
		0,
		0)
}

func GlRectsv(v1 /*const*/ *GLshort, v2 /*const*/ *GLshort) {
	syscall3(glRectsv, 2,
		uintptr(unsafe.Pointer(v1)),
		uintptr(unsafe.Pointer(v2)),
		0)
}

func GlRenderMode(mode GLenum) GLint {
	ret1 := syscall3(glRenderMode, 1,
		uintptr(mode),
		0,
		0)
	return GLint(ret1)
}

func GlRotated(angle GLdouble, x GLdouble, y GLdouble, z GLdouble) {
	syscall6(glRotated, 4,
		uintptr(angle),
		uintptr(x),
		uintptr(y),
		uintptr(z),
		0,
		0)
}

func GlRotatef(angle GLfloat, x GLfloat, y GLfloat, z GLfloat) {
	syscall6(glRotatef, 4,
		uintptr(angle),
		uintptr(x),
		uintptr(y),
		uintptr(z),
		0,
		0)
}

func GlScaled(x GLdouble, y GLdouble, z GLdouble) {
	syscall3(glScaled, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlScalef(x GLfloat, y GLfloat, z GLfloat) {
	syscall3(glScalef, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlScissor(x GLint, y GLint, width GLsizei, height GLsizei) {
	syscall6(glScissor, 4,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		0,
		0)
}

func GlSelectBuffer(size GLsizei, buffer *GLuint) {
	syscall3(glSelectBuffer, 2,
		uintptr(size),
		uintptr(unsafe.Pointer(buffer)),
		0)
}

func GlShadeModel(mode GLenum) {
	syscall3(glShadeModel, 1,
		uintptr(mode),
		0,
		0)
}

func GlStencilFunc(aFunc GLenum, ref GLint, mask GLuint) {
	syscall3(glStencilFunc, 3,
		uintptr(aFunc),
		uintptr(ref),
		uintptr(mask))
}

func GlStencilMask(mask GLuint) {
	syscall3(glStencilMask, 1,
		uintptr(mask),
		0,
		0)
}

func GlStencilOp(fail GLenum, zfail GLenum, zpass GLenum) {
	syscall3(glStencilOp, 3,
		uintptr(fail),
		uintptr(zfail),
		uintptr(zpass))
}

func GlTexCoord1d(s GLdouble) {
	syscall3(glTexCoord1d, 1,
		uintptr(s),
		0,
		0)
}

func GlTexCoord1dv(v /*const*/ *GLdouble) {
	syscall3(glTexCoord1dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord1f(s GLfloat) {
	syscall3(glTexCoord1f, 1,
		uintptr(s),
		0,
		0)
}

func GlTexCoord1fv(v /*const*/ *GLfloat) {
	syscall3(glTexCoord1fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord1i(s GLint) {
	syscall3(glTexCoord1i, 1,
		uintptr(s),
		0,
		0)
}

func GlTexCoord1iv(v /*const*/ *GLint) {
	syscall3(glTexCoord1iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord1s(s GLshort) {
	syscall3(glTexCoord1s, 1,
		uintptr(s),
		0,
		0)
}

func GlTexCoord1sv(v /*const*/ *GLshort) {
	syscall3(glTexCoord1sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord2d(s GLdouble, t GLdouble) {
	syscall3(glTexCoord2d, 2,
		uintptr(s),
		uintptr(t),
		0)
}

func GlTexCoord2dv(v /*const*/ *GLdouble) {
	syscall3(glTexCoord2dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord2f(s GLfloat, t GLfloat) {
	syscall3(glTexCoord2f, 2,
		uintptr(s),
		uintptr(t),
		0)
}

func GlTexCoord2fv(v /*const*/ *GLfloat) {
	syscall3(glTexCoord2fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord2i(s GLint, t GLint) {
	syscall3(glTexCoord2i, 2,
		uintptr(s),
		uintptr(t),
		0)
}

func GlTexCoord2iv(v /*const*/ *GLint) {
	syscall3(glTexCoord2iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord2s(s GLshort, t GLshort) {
	syscall3(glTexCoord2s, 2,
		uintptr(s),
		uintptr(t),
		0)
}

func GlTexCoord2sv(v /*const*/ *GLshort) {
	syscall3(glTexCoord2sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord3d(s GLdouble, t GLdouble, r GLdouble) {
	syscall3(glTexCoord3d, 3,
		uintptr(s),
		uintptr(t),
		uintptr(r))
}

func GlTexCoord3dv(v /*const*/ *GLdouble) {
	syscall3(glTexCoord3dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord3f(s GLfloat, t GLfloat, r GLfloat) {
	syscall3(glTexCoord3f, 3,
		uintptr(s),
		uintptr(t),
		uintptr(r))
}

func GlTexCoord3fv(v /*const*/ *GLfloat) {
	syscall3(glTexCoord3fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord3i(s GLint, t GLint, r GLint) {
	syscall3(glTexCoord3i, 3,
		uintptr(s),
		uintptr(t),
		uintptr(r))
}

func GlTexCoord3iv(v /*const*/ *GLint) {
	syscall3(glTexCoord3iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord3s(s GLshort, t GLshort, r GLshort) {
	syscall3(glTexCoord3s, 3,
		uintptr(s),
		uintptr(t),
		uintptr(r))
}

func GlTexCoord3sv(v /*const*/ *GLshort) {
	syscall3(glTexCoord3sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord4d(s GLdouble, t GLdouble, r GLdouble, q GLdouble) {
	syscall6(glTexCoord4d, 4,
		uintptr(s),
		uintptr(t),
		uintptr(r),
		uintptr(q),
		0,
		0)
}

func GlTexCoord4dv(v /*const*/ *GLdouble) {
	syscall3(glTexCoord4dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord4f(s GLfloat, t GLfloat, r GLfloat, q GLfloat) {
	syscall6(glTexCoord4f, 4,
		uintptr(s),
		uintptr(t),
		uintptr(r),
		uintptr(q),
		0,
		0)
}

func GlTexCoord4fv(v /*const*/ *GLfloat) {
	syscall3(glTexCoord4fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord4i(s GLint, t GLint, r GLint, q GLint) {
	syscall6(glTexCoord4i, 4,
		uintptr(s),
		uintptr(t),
		uintptr(r),
		uintptr(q),
		0,
		0)
}

func GlTexCoord4iv(v /*const*/ *GLint) {
	syscall3(glTexCoord4iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoord4s(s GLshort, t GLshort, r GLshort, q GLshort) {
	syscall6(glTexCoord4s, 4,
		uintptr(s),
		uintptr(t),
		uintptr(r),
		uintptr(q),
		0,
		0)
}

func GlTexCoord4sv(v /*const*/ *GLshort) {
	syscall3(glTexCoord4sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlTexCoordPointer(size GLint, aType GLenum, stride GLsizei, pointer /*const*/ uintptr) {
	syscall6(glTexCoordPointer, 4,
		uintptr(size),
		uintptr(aType),
		uintptr(stride),
		pointer,
		0,
		0)
}

func GlTexEnvf(target GLenum, pname GLenum, param GLfloat) {
	syscall3(glTexEnvf, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(param))
}

func GlTexEnvfv(target GLenum, pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glTexEnvfv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexEnvi(target GLenum, pname GLenum, param GLint) {
	syscall3(glTexEnvi, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(param))
}

func GlTexEnviv(target GLenum, pname GLenum, params /*const*/ *GLint) {
	syscall3(glTexEnviv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexGend(coord GLenum, pname GLenum, param GLdouble) {
	syscall3(glTexGend, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(param))
}

func GlTexGendv(coord GLenum, pname GLenum, params /*const*/ *GLdouble) {
	syscall3(glTexGendv, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexGenf(coord GLenum, pname GLenum, param GLfloat) {
	syscall3(glTexGenf, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(param))
}

func GlTexGenfv(coord GLenum, pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glTexGenfv, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexGeni(coord GLenum, pname GLenum, param GLint) {
	syscall3(glTexGeni, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(param))
}

func GlTexGeniv(coord GLenum, pname GLenum, params /*const*/ *GLint) {
	syscall3(glTexGeniv, 3,
		uintptr(coord),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexImage1D(target GLenum, level GLint, internalformat GLint, width GLsizei, border GLint, format GLenum, aType GLenum, pixels /*const*/ uintptr) {
	syscall9(glTexImage1D, 8,
		uintptr(target),
		uintptr(level),
		uintptr(internalformat),
		uintptr(width),
		uintptr(border),
		uintptr(format),
		uintptr(aType),
		pixels,
		0)
}

func GlTexImage2D(target GLenum, level GLint, internalformat GLint, width GLsizei, height GLsizei, border GLint, format GLenum, aType GLenum, pixels /*const*/ uintptr) {
	syscall9(glTexImage2D, 9,
		uintptr(target),
		uintptr(level),
		uintptr(internalformat),
		uintptr(width),
		uintptr(height),
		uintptr(border),
		uintptr(format),
		uintptr(aType),
		pixels)
}

func GlTexParameterf(target GLenum, pname GLenum, param GLfloat) {
	syscall3(glTexParameterf, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(param))
}

func GlTexParameterfv(target GLenum, pname GLenum, params /*const*/ *GLfloat) {
	syscall3(glTexParameterfv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexParameteri(target GLenum, pname GLenum, param GLint) {
	syscall3(glTexParameteri, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(param))
}

func GlTexParameteriv(target GLenum, pname GLenum, params /*const*/ *GLint) {
	syscall3(glTexParameteriv, 3,
		uintptr(target),
		uintptr(pname),
		uintptr(unsafe.Pointer(params)))
}

func GlTexSubImage1D(target GLenum, level GLint, xoffset GLint, width GLsizei, format GLenum, aType GLenum, pixels /*const*/ uintptr) {
	syscall9(glTexSubImage1D, 7,
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(width),
		uintptr(format),
		uintptr(aType),
		pixels,
		0,
		0)
}

func GlTexSubImage2D(target GLenum, level GLint, xoffset GLint, yoffset GLint, width GLsizei, height GLsizei, format GLenum, aType GLenum, pixels /*const*/ uintptr) {
	syscall9(glTexSubImage2D, 9,
		uintptr(target),
		uintptr(level),
		uintptr(xoffset),
		uintptr(yoffset),
		uintptr(width),
		uintptr(height),
		uintptr(format),
		uintptr(aType),
		pixels)
}

func GlTranslated(x GLdouble, y GLdouble, z GLdouble) {
	syscall3(glTranslated, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlTranslatef(x GLfloat, y GLfloat, z GLfloat) {
	syscall3(glTranslatef, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlVertex2d(x GLdouble, y GLdouble) {
	syscall3(glVertex2d, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlVertex2dv(v /*const*/ *GLdouble) {
	syscall3(glVertex2dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex2f(x GLfloat, y GLfloat) {
	syscall3(glVertex2f, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlVertex2fv(v /*const*/ *GLfloat) {
	syscall3(glVertex2fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex2i(x GLint, y GLint) {
	syscall3(glVertex2i, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlVertex2iv(v /*const*/ *GLint) {
	syscall3(glVertex2iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex2s(x GLshort, y GLshort) {
	syscall3(glVertex2s, 2,
		uintptr(x),
		uintptr(y),
		0)
}

func GlVertex2sv(v /*const*/ *GLshort) {
	syscall3(glVertex2sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex3d(x GLdouble, y GLdouble, z GLdouble) {
	syscall3(glVertex3d, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlVertex3dv(v /*const*/ *GLdouble) {
	syscall3(glVertex3dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex3f(x GLfloat, y GLfloat, z GLfloat) {
	syscall3(glVertex3f, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlVertex3fv(v /*const*/ *GLfloat) {
	syscall3(glVertex3fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex3i(x GLint, y GLint, z GLint) {
	syscall3(glVertex3i, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlVertex3iv(v /*const*/ *GLint) {
	syscall3(glVertex3iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex3s(x GLshort, y GLshort, z GLshort) {
	syscall3(glVertex3s, 3,
		uintptr(x),
		uintptr(y),
		uintptr(z))
}

func GlVertex3sv(v /*const*/ *GLshort) {
	syscall3(glVertex3sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex4d(x GLdouble, y GLdouble, z GLdouble, w GLdouble) {
	syscall6(glVertex4d, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlVertex4dv(v /*const*/ *GLdouble) {
	syscall3(glVertex4dv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex4f(x GLfloat, y GLfloat, z GLfloat, w GLfloat) {
	syscall6(glVertex4f, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlVertex4fv(v /*const*/ *GLfloat) {
	syscall3(glVertex4fv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex4i(x GLint, y GLint, z GLint, w GLint) {
	syscall6(glVertex4i, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlVertex4iv(v /*const*/ *GLint) {
	syscall3(glVertex4iv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertex4s(x GLshort, y GLshort, z GLshort, w GLshort) {
	syscall6(glVertex4s, 4,
		uintptr(x),
		uintptr(y),
		uintptr(z),
		uintptr(w),
		0,
		0)
}

func GlVertex4sv(v /*const*/ *GLshort) {
	syscall3(glVertex4sv, 1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)
}

func GlVertexPointer(size GLint, aType GLenum, stride GLsizei, pointer /*const*/ uintptr) {
	syscall6(glVertexPointer, 4,
		uintptr(size),
		uintptr(aType),
		uintptr(stride),
		pointer,
		0,
		0)
}

func GlViewport(x GLint, y GLint, width GLsizei, height GLsizei) {
	syscall6(glViewport, 4,
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		0,
		0)
}

func WglCopyContext(unnamed0 HGLRC, unnamed1 HGLRC, unnamed2 uint32) bool {
	ret1 := syscall3(wglCopyContext, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2))
	return ret1 != 0
}

func WglCreateContext(unnamed0 HDC) HGLRC {
	ret1 := syscall3(wglCreateContext, 1,
		uintptr(unnamed0),
		0,
		0)
	return HGLRC(ret1)
}

func WglCreateLayerContext(unnamed0 HDC, unnamed1 int32) HGLRC {
	ret1 := syscall3(wglCreateLayerContext, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return HGLRC(ret1)
}

func WglDeleteContext(unnamed0 HGLRC) bool {
	ret1 := syscall3(wglDeleteContext, 1,
		uintptr(unnamed0),
		0,
		0)
	return ret1 != 0
}

func WglDescribeLayerPlane(unnamed0 HDC, unnamed1 int32, unnamed2 int32, unnamed3 uint32, unnamed4 *LAYERPLANEDESCRIPTOR) bool {
	ret1 := syscall6(wglDescribeLayerPlane, 5,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		uintptr(unsafe.Pointer(unnamed4)),
		0)
	return ret1 != 0
}

func WglGetCurrentContext() HGLRC {
	ret1 := syscall3(wglGetCurrentContext, 0,
		0,
		0,
		0)
	return HGLRC(ret1)
}

func WglGetCurrentDC() HDC {
	ret1 := syscall3(wglGetCurrentDC, 0,
		0,
		0,
		0)
	return HDC(ret1)
}

func WglGetLayerPaletteEntries(unnamed0 HDC, unnamed1 int32, unnamed2 int32, unnamed3 int32, unnamed4 *COLORREF) int32 {
	ret1 := syscall6(wglGetLayerPaletteEntries, 5,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		uintptr(unsafe.Pointer(unnamed4)),
		0)
	return int32(ret1)
}

func WglGetProcAddress(unnamed0 /*const*/ LPCSTR) PROC {
	ret1 := syscall3(wglGetProcAddress, 1,
		uintptr(unsafe.Pointer(unnamed0)),
		0,
		0)
	return PROC(ret1)
}

func WglMakeCurrent(unnamed0 HDC, unnamed1 HGLRC) bool {
	ret1 := syscall3(wglMakeCurrent, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func WglRealizeLayerPalette(unnamed0 HDC, unnamed1 int32, unnamed2 bool) bool {
	ret1 := syscall3(wglRealizeLayerPalette, 3,
		uintptr(unnamed0),
		uintptr(unnamed1),
		getUintptrFromBool(unnamed2))
	return ret1 != 0
}

func WglSetLayerPaletteEntries(unnamed0 HDC, unnamed1 int32, unnamed2 int32, unnamed3 int32, unnamed4 /*const*/ *COLORREF) int32 {
	ret1 := syscall6(wglSetLayerPaletteEntries, 5,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		uintptr(unsafe.Pointer(unnamed4)),
		0)
	return int32(ret1)
}

func WglShareLists(unnamed0 HGLRC, unnamed1 HGLRC) bool {
	ret1 := syscall3(wglShareLists, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func WglSwapLayerBuffers(unnamed0 HDC, unnamed1 uint32) bool {
	ret1 := syscall3(wglSwapLayerBuffers, 2,
		uintptr(unnamed0),
		uintptr(unnamed1),
		0)
	return ret1 != 0
}

func WglSwapMultipleBuffers(unnamed0 uint32, unnamed1 /*const*/ *WGLSWAP) uint32 {
	ret1 := syscall3(wglSwapMultipleBuffers, 2,
		uintptr(unnamed0),
		uintptr(unsafe.Pointer(unnamed1)),
		0)
	return uint32(ret1)
}

func WglUseFontBitmaps(unnamed0 HDC, unnamed1 uint32, unnamed2 uint32, unnamed3 uint32) bool {
	ret1 := syscall6(wglUseFontBitmaps, 4,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		0,
		0)
	return ret1 != 0
}

func WglUseFontOutlines(unnamed0 HDC, unnamed1 uint32, unnamed2 uint32, unnamed3 uint32, unnamed4 float32, unnamed5 float32, unnamed6 int32, unnamed7 *GLYPHMETRICSFLOAT) bool {
	ret1 := syscall9(wglUseFontOutlines, 8,
		uintptr(unnamed0),
		uintptr(unnamed1),
		uintptr(unnamed2),
		uintptr(unnamed3),
		uintptr(unnamed4),
		uintptr(unnamed5),
		uintptr(unnamed6),
		uintptr(unsafe.Pointer(unnamed7)),
		0)
	return ret1 != 0
}

func GlDebugEntry(unknown1 GLint, unknown2 GLint) GLint {
	ret1 := syscall3(glDebugEntry, 2,
		uintptr(unknown1),
		uintptr(unknown2),
		0)
	return GLint(ret1)
}

func WglChoosePixelFormat(hdc HDC, ppfd /*const*/ *PIXELFORMATDESCRIPTOR) int32 {
	ret1 := syscall3(wglChoosePixelFormat, 2,
		uintptr(hdc),
		uintptr(unsafe.Pointer(ppfd)),
		0)
	return int32(ret1)
}

func WglDescribePixelFormat(hdc HDC, format int32, size uint32, descr *PIXELFORMATDESCRIPTOR) int32 {
	ret1 := syscall6(wglDescribePixelFormat, 4,
		uintptr(hdc),
		uintptr(format),
		uintptr(size),
		uintptr(unsafe.Pointer(descr)),
		0,
		0)
	return int32(ret1)
}

func WglGetPixelFormat(hdc HDC) int32 {
	ret1 := syscall3(wglGetPixelFormat, 1,
		uintptr(hdc),
		0,
		0)
	return int32(ret1)
}

func WglSetPixelFormat(hdc HDC, format int32, descr /*const*/ *PIXELFORMATDESCRIPTOR) bool {
	ret1 := syscall3(wglSetPixelFormat, 3,
		uintptr(hdc),
		uintptr(format),
		uintptr(unsafe.Pointer(descr)))
	return ret1 != 0
}
