// Copyright (C) 2025 The Qt Company Ltd.
// SPDX-License-Identifier: MIT

#include "container_types.h"
#include "core_types.h"
#include "gui_types.h"
#include "network_types.h"
#include "variant_types.h"
#include "widgets_types.h"

#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication app(argc, argv);

    auto coreTypes = CoreTypes();

    auto containerTypes = ContainerTypes();

    auto variantTypes = VariantTypes();

    auto networkTypes = NetworkTypes();

    auto guiTypes = GuiTypes();

    auto widgetTypes = WidgetsTypes();

    return 0;
}
