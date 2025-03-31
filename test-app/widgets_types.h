// Copyright (C) 2025 The Qt Company Ltd.
// SPDX-License-Identifier: MIT

#include <QtWidgets>

class WidgetsTypes : public QObject
{
    Q_OBJECT

public:
    using QObject::QObject;

    QSizePolicy qSizePolicy = QSizePolicy(QSizePolicy::Expanding, QSizePolicy::Expanding);
    QWidget qWidget;
};
