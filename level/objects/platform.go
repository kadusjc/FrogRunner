components {
  id: "platform"
  component: "/level/objects/platform.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/level/level.atlas\"\n"
  "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "danger_shapes"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"danger\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -199.0\n"
  "      y: -1.0\n"
  "      z: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 198.0\n"
  "      y: 4.0\n"
  "      z: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -9.0\n"
  "      y: -77.0\n"
  "      z: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: -0.7075842\n"
  "      w: 0.70662904\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  data: 22.855\n"
  "  data: 52.9435\n"
  "  data: 9.8\n"
  "  data: 22.855\n"
  "  data: 52.9435\n"
  "  data: 9.8\n"
  "  data: 22.855\n"
  "  data: 165.494\n"
  "  data: 9.8\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"geometry\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "      y: 60.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 175.405\n"
  "  data: 12.2925\n"
  "  data: 11.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "tile_set: \"/level/level.atlas\"\n"
  "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -197.0
    y: -3.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite2"
  type: "sprite"
  data: "tile_set: \"/level/level.atlas\"\n"
  "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 192.0
    y: 3.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: -0.99999934
    w: 0.0011224704
  }
}
embedded_components {
  id: "sprite3"
  type: "sprite"
  data: "tile_set: \"/level/level.atlas\"\n"
  "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: -115.0
    y: -82.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite4"
  type: "sprite"
  data: "tile_set: \"/level/level.atlas\"\n"
  "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 13.0
    y: -82.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "sprite5"
  type: "sprite"
  data: "tile_set: \"/level/level.atlas\"\n"
  "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 109.0
    y: -80.0
    z: -0.1
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.70710677
    w: 0.70710677
  }
}
