package library

const (
	SPA_HP                                 = 0
	SPA_AC                                 = 1
	SPA_ATTACK_POWER                       = 2
	SPA_MOVEMENT_RATE                      = 3
	SPA_STR                                = 4
	SPA_DEX                                = 5
	SPA_AGI                                = 6
	SPA_STA                                = 7
	SPA_INT                                = 8
	SPA_WIS                                = 9
	SPA_CHA                                = 10
	SPA_HASTE                              = 11 // "Melee Speed"
	SPA_INVISIBILITY                       = 12
	SPA_SEE_INVIS                          = 13
	SPA_ENDURING_BREATH                    = 14
	SPA_MANA                               = 15
	SPA_NPC_FRENZY                         = 16
	SPA_NPC_AWARENESS                      = 17
	SPA_NPC_AGGRO                          = 18
	SPA_NPC_FACTION                        = 19
	SPA_BLINDNESS                          = 20
	SPA_STUN                               = 21
	SPA_CHARM                              = 22
	SPA_FEAR                               = 23
	SPA_FATIGUE                            = 24
	SPA_BIND_AFFINITY                      = 25
	SPA_GATE                               = 26
	SPA_DISPEL_MAGIC                       = 27
	SPA_INVIS_VS_UNDEAD                    = 28
	SPA_INVIS_VS_ANIMALS                   = 29
	SPA_NPC_AGGRO_RADIUS                   = 30
	SPA_ENTHRALL                           = 31
	SPA_CREATE_ITEM                        = 32
	SPA_SUMMON_PET                         = 33
	SPA_CONFUSE                            = 34
	SPA_DISEASE                            = 35
	SPA_POISON                             = 36
	SPA_DETECT_HOSTILE                     = 37
	SPA_DETECT_MAGIC                       = 38
	SPA_NO_TWINCAST                        = 39
	SPA_INVULNERABILITY                    = 40
	SPA_BANISH                             = 41
	SPA_SHADOW_STEP                        = 42
	SPA_BERSERK                            = 43
	SPA_LYCANTHROPY                        = 44
	SPA_VAMPIRISM                          = 45
	SPA_RESIST_FIRE                        = 46
	SPA_RESIST_COLD                        = 47
	SPA_RESIST_POISON                      = 48
	SPA_RESIST_DISEASE                     = 49
	SPA_RESIST_MAGIC                       = 50
	SPA_DETECT_TRAPS                       = 51
	SPA_DETECT_UNDEAD                      = 52
	SPA_DETECT_SUMMONED                    = 53
	SPA_DETECT_ANIMALS                     = 54
	SPA_STONESKIN                          = 55
	SPA_TRUE_NORTH                         = 56
	SPA_LEVITATION                         = 57
	SPA_CHANGE_FORM                        = 58
	SPA_DAMAGE_SHIELD                      = 59
	SPA_TRANSFER_ITEM                      = 60
	SPA_ITEM_LORE                          = 61
	SPA_ITEM_IDENTIFY                      = 62
	SPA_NPC_WIPE_HATE_LIST                 = 63
	SPA_SPIN_STUN                          = 64
	SPA_INFRAVISION                        = 65
	SPA_ULTRAVISION                        = 66
	SPA_EYE_OF_ZOMM                        = 67
	SPA_RECLAIM_ENERGY                     = 68
	SPA_MAX_HP                             = 69
	SPA_CORPSE_BOMB                        = 70
	SPA_CREATE_UNDEAD                      = 71
	SPA_PRESERVE_CORPSE                    = 72
	SPA_BIND_SIGHT                         = 73
	SPA_FEIGN_DEATH                        = 74
	SPA_VENTRILOQUISM                      = 75
	SPA_SENTINEL                           = 76
	SPA_LOCATE_CORPSE                      = 77
	SPA_SPELL_SHIELD                       = 78
	SPA_INSTANT_HP                         = 79
	SPA_ENCHANT_LIGHT                      = 80
	SPA_RESURRECT                          = 81
	SPA_SUMMON_TARGET                      = 82
	SPA_PORTAL                             = 83
	SPA_HP_NPC_ONLY                        = 84
	SPA_MELEE_PROC                         = 85
	SPA_NPC_HELP_RADIUS                    = 86
	SPA_MAGNIFICATION                      = 87
	SPA_EVACUATE                           = 88
	SPA_HEIGHT                             = 89
	SPA_IGNORE_PET                         = 90
	SPA_SUMMON_CORPSE                      = 91
	SPA_HATE                               = 92
	SPA_WEATHER_CONTROL                    = 93
	SPA_FRAGILE                            = 94
	SPA_SACRIFICE                          = 95
	SPA_SILENCE                            = 96
	SPA_MAX_MANA                           = 97
	SPA_BARD_HASTE                         = 98
	SPA_ROOT                               = 99
	SPA_HEALDOT                            = 100
	SPA_COMPLETEHEAL                       = 101
	SPA_PET_FEARLESS                       = 102
	SPA_CALL_PET                           = 103
	SPA_TRANSLOCATE                        = 104
	SPA_NPC_ANTI_GATE                      = 105
	SPA_BEASTLORD_PET                      = 106
	SPA_ALTER_PET_LEVEL                    = 107
	SPA_FAMILIAR                           = 108
	SPA_CREATE_ITEM_IN_BAG                 = 109
	SPA_ARCHERY                            = 110
	SPA_RESIST_ALL                         = 111
	SPA_FIZZLE_SKILL                       = 112
	SPA_SUMMON_MOUNT                       = 113
	SPA_MODIFY_HATE                        = 114
	SPA_CORNUCOPIA                         = 115
	SPA_CURSE                              = 116
	SPA_HIT_MAGIC                          = 117
	SPA_AMPLIFICATION                      = 118
	SPA_ATTACK_SPEED_MAX                   = 119
	SPA_HEALMOD                            = 120
	SPA_IRONMAIDEN                         = 121
	SPA_REDUCESKILL                        = 122
	SPA_IMMUNITY                           = 123
	SPA_FOCUS_DAMAGE_MOD                   = 124
	SPA_FOCUS_HEAL_MOD                     = 125
	SPA_FOCUS_RESIST_MOD                   = 126
	SPA_FOCUS_CAST_TIME_MOD                = 127
	SPA_FOCUS_DURATION_MOD                 = 128
	SPA_FOCUS_RANGE_MOD                    = 129
	SPA_FOCUS_HATE_MOD                     = 130
	SPA_FOCUS_REAGENT_MOD                  = 131
	SPA_FOCUS_MANACOST_MOD                 = 132
	SPA_FOCUS_STUNTIME_MOD                 = 133
	SPA_FOCUS_LEVEL_MAX                    = 134
	SPA_FOCUS_RESIST_TYPE                  = 135
	SPA_FOCUS_TARGET_TYPE                  = 136
	SPA_FOCUS_WHICH_SPA                    = 137
	SPA_FOCUS_BENEFICIAL                   = 138
	SPA_FOCUS_WHICH_SPELL                  = 139
	SPA_FOCUS_DURATION_MIN                 = 140
	SPA_FOCUS_INSTANT_ONLY                 = 141
	SPA_FOCUS_LEVEL_MIN                    = 142
	SPA_FOCUS_CASTTIME_MIN                 = 143
	SPA_FOCUS_CASTTIME_MAX                 = 144
	SPA_NPC_PORTAL_WARDER_BANISH           = 145
	SPA_PORTAL_LOCATIONS                   = 146
	SPA_PERCENT_HEAL                       = 147
	SPA_STACKING_BLOCK                     = 148
	SPA_STRIP_VIRTUAL_SLOT                 = 149
	SPA_DIVINE_INTERVENTION                = 150
	SPA_POCKET_PET                         = 151
	SPA_PET_SWARM                          = 152
	SPA_HEALTH_BALANCE                     = 153
	SPA_CANCEL_NEGATIVE_MAGIC              = 154
	SPA_POP_RESURRECT                      = 155
	SPA_MIRROR                             = 156
	SPA_FEEDBACK                           = 157
	SPA_REFLECT                            = 158
	SPA_MODIFY_ALL_STATS                   = 159
	SPA_CHANGE_SOBRIETY                    = 160
	SPA_SPELL_GUARD                        = 161
	SPA_MELEE_GUARD                        = 162
	SPA_ABSORB_HIT                         = 163
	SPA_OBJECT_SENSE_TRAP                  = 164
	SPA_OBJECT_DISARM_TRAP                 = 165
	SPA_OBJECT_PICKLOCK                    = 166
	SPA_FOCUS_PET                          = 167
	SPA_DEFENSIVE                          = 168
	SPA_CRITICAL_MELEE                     = 169
	SPA_CRITICAL_SPELL                     = 170
	SPA_CRIPPLING_BLOW                     = 171
	SPA_EVASION                            = 172
	SPA_RIPOSTE                            = 173
	SPA_DODGE                              = 174
	SPA_PARRY                              = 175
	SPA_DUAL_WIELD                         = 176
	SPA_DOUBLE_ATTACK                      = 177
	SPA_MELEE_LIFETAP                      = 178
	SPA_PURETONE                           = 179
	SPA_SANCTIFICATION                     = 180
	SPA_FEARLESS                           = 181
	SPA_HUNDRED_HANDS                      = 182
	SPA_SKILL_INCREASE_CHANCE              = 183 // Unused
	SPA_ACCURACY                           = 184
	SPA_SKILL_DAMAGE_MOD                   = 185
	SPA_MIN_DAMAGE_DONE_MOD                = 186
	SPA_MANA_BALANCE                       = 187
	SPA_BLOCK                              = 188
	SPA_ENDURANCE                          = 189
	SPA_INCREASE_MAX_ENDURANCE             = 190
	SPA_AMNESIA                            = 191
	SPA_HATE_OVER_TIME                     = 192
	SPA_SKILL_ATTACK                       = 193
	SPA_FADE                               = 194
	SPA_STUN_RESIST                        = 195
	SPA_STRIKETHROUGH1                     = 196 // Deprecated
	SPA_SKILL_DAMAGE_TAKEN                 = 197
	SPA_INSTANT_ENDURANCE                  = 198
	SPA_TAUNT                              = 199
	SPA_PROC_CHANCE                        = 200
	SPA_RANGE_ABILITY                      = 201
	SPA_ILLUSION_OTHERS                    = 202
	SPA_MASS_GROUP_BUFF                    = 203
	SPA_GROUP_FEAR_IMMUNITY                = 204
	SPA_RAMPAGE                            = 205
	SPA_AE_TAUNT                           = 206
	SPA_FLESH_TO_BONE                      = 207
	SPA_PURGE_POISON                       = 208
	SPA_CANCEL_BENEFICIAL                  = 209
	SPA_SHIELD_CASTER                      = 210
	SPA_DESTRUCTIVE_FORCE                  = 211
	SPA_FOCUS_FRENZIED_DEVASTATION         = 212
	SPA_PET_PCT_MAX_HP                     = 213
	SPA_HP_MAX_HP                          = 214
	SPA_PET_PCT_AVOIDANCE                  = 215
	SPA_MELEE_ACCURACY                     = 216
	SPA_HEADSHOT                           = 217
	SPA_PET_CRIT_MELEE                     = 218
	SPA_SLAY_UNDEAD                        = 219
	SPA_INCREASE_SKILL_DAMAGE              = 220
	SPA_REDUCE_WEIGHT                      = 221
	SPA_BLOCK_BEHIND                       = 222
	SPA_DOUBLE_RIPOSTE                     = 223
	SPA_ADD_RIPOSTE                        = 224
	SPA_GIVE_DOUBLE_ATTACK                 = 225
	SPA_2H_BASH                            = 226
	SPA_REDUCE_SKILL_TIMER                 = 227
	SPA_ACROBATICS                         = 228
	SPA_CAST_THROUGH_STUN                  = 229
	SPA_EXTENDED_SHIELDING                 = 230
	SPA_BASH_CHANCE                        = 231
	SPA_DIVINE_SAVE                        = 232
	SPA_METABOLISM                         = 233
	SPA_POISON_MASTERY                     = 234
	SPA_FOCUS_CHANNELING                   = 235
	SPA_FREE_PET                           = 236
	SPA_PET_AFFINITY                       = 237
	SPA_PERM_ILLUSION                      = 238
	SPA_STONEWALL                          = 239
	SPA_STRING_UNBREAKABLE                 = 240
	SPA_IMPROVE_RECLAIM_ENERGY             = 241
	SPA_INCREASE_CHANGE_MEMWIPE            = 242
	SPA_ENHANCED_CHARM                     = 243
	SPA_ENHANCED_ROOT                      = 244
	SPA_TRAP_CIRCUMVENTION                 = 245
	SPA_INCREASE_AIR_SUPPLY                = 246
	SPA_INCREASE_MAX_SKILL                 = 247
	SPA_EXTRA_SPECIALIZATION               = 248
	SPA_OFFHAND_MIN_WEAPON_DAMAGE          = 249
	SPA_INCREASE_PROC_CHANCE               = 250
	SPA_ENDLESS_QUIVER                     = 251
	SPA_BACKSTAB_FRONT                     = 252
	SPA_CHAOTIC_STAB                       = 253
	SPA_NOSPELL                            = 254
	SPA_SHIELDING_DURATION_MOD             = 255
	SPA_SHROUD_OF_STEALTH                  = 256
	SPA_GIVE_PET_HOLD                      = 257 // Deprecated
	SPA_TRIPLE_BACKSTAB                    = 258
	SPA_AC_LIMIT_MOD                       = 259
	SPA_ADD_INSTRUMENT_MOD                 = 260
	SPA_SONG_MOD_CAP                       = 261
	SPA_INCREASE_STAT_CAP                  = 262
	SPA_TRADESKILL_MASTERY                 = 263
	SPA_REDUCE_AA_TIMER                    = 264
	SPA_NO_FIZZLE                          = 265
	SPA_ADD_2H_ATTACK_CHANCE               = 266
	SPA_ADD_PET_COMMANDS                   = 267
	SPA_ALCHEMY_FAIL_RATE                  = 268
	SPA_FIRST_AID                          = 269
	SPA_EXTEND_SONG_RANGE                  = 270
	SPA_BASE_RUN_MOD                       = 271
	SPA_INCREASE_CASTING_LEVEL             = 272
	SPA_DOTCRIT                            = 273
	SPA_HEALCRIT                           = 274
	SPA_MENDCRIT                           = 275
	SPA_DUAL_WIELD_AMT                     = 276
	SPA_EXTRA_DI_CHANCE                    = 277
	SPA_FINISHING_BLOW                     = 278
	SPA_FLURRY                             = 279
	SPA_PET_FLURRY                         = 280
	SPA_PET_FEIGN                          = 281
	SPA_INCREASE_BANDAGE_AMT               = 282
	SPA_WU_ATTACK                          = 283
	SPA_IMPROVE_LOH                        = 284
	SPA_NIMBLE_EVASION                     = 285
	SPA_FOCUS_DAMAGE_AMT                   = 286
	SPA_FOCUS_DURATION_AMT                 = 287
	SPA_ADD_PROC_HIT                       = 288
	SPA_DOOM_EFFECT                        = 289
	SPA_INCREASE_RUN_SPEED_CAP             = 290
	SPA_PURIFY                             = 291
	SPA_STRIKETHROUGH                      = 292
	SPA_STUN_RESIST2                       = 293
	SPA_SPELL_CRIT_CHANCE                  = 294
	SPA_REDUCE_SPECIAL_TIMER               = 295
	SPA_FOCUS_DAMAGE_MOD_DETRIMENTAL       = 296
	SPA_FOCUS_DAMAGE_AMT_DETRIMENTAL       = 297
	SPA_TINY_COMPANION                     = 298
	SPA_WAKE_DEAD                          = 299
	SPA_DOPPELGANGER                       = 300
	SPA_INCREASE_RANGE_DMG                 = 301
	SPA_FOCUS_DAMAGE_MOD_CRIT              = 302
	SPA_FOCUS_DAMAGE_AMT_CRIT              = 303
	SPA_SECONDARY_RIPOSTE_MOD              = 304
	SPA_DAMAGE_SHIELD_MOD                  = 305
	SPA_WEAK_DEAD_2                        = 306
	SPA_APPRAISAL                          = 307
	SPA_ZONE_SUSPEND_MINION                = 308
	SPA_TELEPORT_CASTERS_BINDPOINT         = 309
	SPA_FOCUS_REUSE_TIMER                  = 310
	SPA_FOCUS_COMBAT_SKILL                 = 311
	SPA_OBSERVER                           = 312
	SPA_FORAGE_MASTER                      = 313
	SPA_IMPROVED_INVIS                     = 314
	SPA_IMPROVED_INVIS_UNDEAD              = 315
	SPA_IMPROVED_INVIS_ANIMALS             = 316
	SPA_INCREASE_WORN_HP_REGEN_CAP         = 317
	SPA_INCREASE_WORN_MANA_REGEN_CAP       = 318
	SPA_CRITICAL_HP_REGEN                  = 319
	SPA_SHIELD_BLOCK_CHANCE                = 320
	SPA_REDUCE_TARGET_HATE                 = 321
	SPA_GATE_STARTING_CITY                 = 322
	SPA_DEFENSIVE_PROC                     = 323
	SPA_HP_FOR_MANA                        = 324
	SPA_NO_BREAK_AE_SNEAK                  = 325
	SPA_ADD_SPELL_SLOTS                    = 326
	SPA_ADD_BUFF_SLOTS                     = 327
	SPA_INCREASE_NEGATIVE_HP_LIMIT         = 328
	SPA_MANA_ABSORB_PCT_DMG                = 329
	SPA_CRIT_ATTACK_MODIFIER               = 330
	SPA_FAIL_ALCHEMY_ITEM_RECOVERY         = 331
	SPA_SUMMON_TO_CORPSE                   = 332
	SPA_DOOM_RUNE_EFFECT                   = 333
	SPA_NO_MOVE_HP                         = 334
	SPA_FOCUSED_IMMUNITY                   = 335
	SPA_ILLUSIONARY_TARGET                 = 336
	SPA_INCREASE_EXP_MOD                   = 337
	SPA_EXPEDIENT_RECOVERY                 = 338
	SPA_FOCUS_CASTING_PROC                 = 339
	SPA_CHANCE_SPELL                       = 340
	SPA_WORN_ATTACK_CAP                    = 341
	SPA_NO_PANIC                           = 342
	SPA_SPELL_INTERRUPT                    = 343
	SPA_ITEM_CHANNELING                    = 344
	SPA_ASSASSINATE_MAX_LEVEL              = 345
	SPA_HEADSHOT_MAX_LEVEL                 = 346
	SPA_DOUBLE_RANGED_ATTACK               = 347
	SPA_FOCUS_MANA_MIN                     = 348
	SPA_INCREASE_SHIELD_DMG                = 349
	SPA_MANABURN                           = 350
	SPA_SPAWN_INTERACTIVE_OBJECT           = 351
	SPA_INCREASE_TRAP_COUNT                = 352
	SPA_INCREASE_SOI_COUNT                 = 353
	SPA_DEACTIVATE_ALL_TRAPS               = 354
	SPA_LEARN_TRAP                         = 355
	SPA_CHANGE_TRIGGER_TYPE                = 356
	SPA_FOCUS_MUTE                         = 357
	SPA_INSTANT_MANA                       = 358
	SPA_PASSIVE_SENSE_TRAP                 = 359
	SPA_PROC_ON_KILL_SHOT                  = 360
	SPA_PROC_ON_DEATH                      = 361
	SPA_POTION_BELT                        = 362
	SPA_BANDOLIER                          = 363
	SPA_ADD_TRIPLE_ATTACK_CHANCE           = 364
	SPA_PROC_ON_SPELL_KILL_SHOT            = 365
	SPA_GROUP_SHIELDING                    = 366
	SPA_MODIFY_BODY_TYPE                   = 367
	SPA_MODIFY_FACTION                     = 368
	SPA_CORRUPTION                         = 369
	SPA_RESIST_CORRUPTION                  = 370
	SPA_SLOW                               = 371
	SPA_GRANT_FORAGING                     = 372
	SPA_DOOM_ALWAYS                        = 373
	SPA_TRIGGER_SPELL                      = 374
	SPA_CRIT_DOT_DMG_MOD                   = 375
	SPA_FLING                              = 376
	SPA_DOOM_ENTITY                        = 377
	SPA_RESIST_OTHER_SPA                   = 378
	SPA_DIRECTIONAL_TELEPORT               = 379
	SPA_EXPLOSIVE_KNOCKBACK                = 380
	SPA_FLING_TOWARD                       = 381
	SPA_SUPPRESSION                        = 382
	SPA_FOCUS_CASTING_PROC_NORMALIZED      = 383
	SPA_FLING_AT                           = 384
	SPA_FOCUS_WHICH_GROUP                  = 385
	SPA_DOOM_DISPELLER                     = 386
	SPA_DOOM_DISPELLEE                     = 387
	SPA_SUMMON_ALL_CORPSES                 = 388
	SPA_REFRESH_SPELL_TIMER                = 389
	SPA_LOCKOUT_SPELL_TIMER                = 390
	SPA_FOCUS_MANA_MAX                     = 391
	SPA_FOCUS_HEAL_AMT                     = 392
	SPA_FOCUS_HEAL_MOD_BENEFICIAL          = 393
	SPA_FOCUS_HEAL_AMT_BENEFICIAL          = 394
	SPA_FOCUS_HEAL_MOD_CRIT                = 395
	SPA_FOCUS_HEAL_AMT_CRIT                = 396
	SPA_ADD_PET_AC                         = 397
	SPA_FOCUS_SWARM_PET_DURATION           = 398
	SPA_FOCUS_TWINCAST_CHANCE              = 399
	SPA_HEALBURN                           = 400
	SPA_MANA_IGNITE                        = 401
	SPA_ENDURANCE_IGNITE                   = 402
	SPA_FOCUS_SPELL_CLASS                  = 403
	SPA_FOCUS_SPELL_SUBCLASS               = 404
	SPA_STAFF_BLOCK_CHANCE                 = 405
	SPA_DOOM_LIMIT_USE                     = 406
	SPA_DOOM_FOCUS_USED                    = 407
	SPA_LIMIT_HP                           = 408
	SPA_LIMIT_MANA                         = 409
	SPA_LIMIT_ENDURANCE                    = 410
	SPA_FOCUS_LIMIT_CLASS                  = 411
	SPA_FOCUS_LIMIT_RACE                   = 412
	SPA_FOCUS_BASE_EFFECTS                 = 413
	SPA_FOCUS_LIMIT_SKILL                  = 414
	SPA_FOCUS_LIMIT_ITEM_CLASS             = 415
	SPA_AC2                                = 416
	SPA_MANA2                              = 417
	SPA_FOCUS_INCREASE_SKILL_DMG_2         = 418
	SPA_PROC_EFFECT_2                      = 419
	SPA_FOCUS_LIMIT_USE                    = 420
	SPA_FOCUS_LIMIT_USE_AMT                = 421
	SPA_FOCUS_LIMIT_USE_MIN                = 422
	SPA_FOCUS_LIMIT_USE_TYPE               = 423
	SPA_GRAVITATE                          = 424
	SPA_FLY                                = 425
	SPA_ADD_EXTENDED_TARGET_SLOTS          = 426
	SPA_SKILL_PROC                         = 427
	SPA_PROC_SKILL_MODIFIER                = 428
	SPA_SKILL_PROC_SUCCESS                 = 429
	SPA_POST_EFFECT                        = 430
	SPA_POST_EFFECT_DATA                   = 431
	SPA_EXPAND_MAX_ACTIVE_TROPHY_BENEFITS  = 432
	SPA_ADD_NORMALIZED_SKILL_MIN_DMG_AMT   = 433
	SPA_ADD_NORMALIZED_SKILL_MIN_DMG_AMT_2 = 434
	SPA_FRAGILE_DEFENSE                    = 435
	SPA_FREEZE_BUFF_TIMER                  = 436
	SPA_TELEPORT_TO_ANCHOR                 = 437
	SPA_TRANSLOCATE_TO_ANCHOR              = 438
	SPA_ASSASSINATE                        = 439
	SPA_FINISHING_BLOW_MAX                 = 440
	SPA_DISTANCE_REMOVAL                   = 441
	SPA_REQUIRE_TARGET_DOOM                = 442
	SPA_REQUIRE_CASTER_DOOM                = 443
	SPA_IMPROVED_TAUNT                     = 444
	SPA_ADD_MERC_SLOT                      = 445
	SPA_STACKER_A                          = 446
	SPA_STACKER_B                          = 447
	SPA_STACKER_C                          = 448
	SPA_STACKER_D                          = 449
	SPA_DOT_GUARD                          = 450
	SPA_MELEE_THRESHOLD_GUARD              = 451
	SPA_SPELL_THRESHOLD_GUARD              = 452
	SPA_MELEE_THRESHOLD_DOOM               = 453
	SPA_SPELL_THRESHOLD_DOOM               = 454
	SPA_ADD_HATE_PCT                       = 455
	SPA_ADD_HATE_OVER_TIME_PCT             = 456
	SPA_RESOURCE_TAP                       = 457
	SPA_FACTION_MOD                        = 458
	SPA_SKILL_DAMAGE_MOD_2                 = 459
	SPA_OVERRIDE_NOT_FOCUSABLE             = 460
	SPA_FOCUS_DAMAGE_MOD_2                 = 461
	SPA_FOCUS_DAMAGE_AMT_2                 = 462
	SPA_SHIELD                             = 463
	SPA_PC_PET_RAMPAGE                     = 464
	SPA_PC_PET_AE_RAMPAGE                  = 465
	SPA_PC_PET_FLURRY                      = 466
	SPA_DAMAGE_SHIELD_MITIGATION_AMT       = 467
	SPA_DAMAGE_SHIELD_MITIGATION_PCT       = 468
	SPA_CHANCE_BEST_IN_SPELL_GROUP         = 469
	SPA_TRIGGER_BEST_IN_SPELL_GROUP        = 470
	SPA_DOUBLE_MELEE_ATTACKS               = 471
	SPA_AA_BUY_NEXT_RANK                   = 472
	SPA_DOUBLE_BACKSTAB_FRONT              = 473
	SPA_PET_MELEE_CRIT_DMG_MOD             = 474
	SPA_TRIGGER_SPELL_NON_ITEM             = 475
	SPA_WEAPON_STANCE                      = 476
	SPA_HATELIST_TO_TOP                    = 477
	SPA_HATELIST_TO_TAIL                   = 478
	SPA_FOCUS_LIMIT_MIN_VALUE              = 479
	SPA_FOCUS_LIMIT_MAX_VALUE              = 480
	SPA_FOCUS_CAST_SPELL_ON_LAND           = 481
	SPA_SKILL_BASE_DAMAGE_MOD              = 482
	SPA_FOCUS_INCOMING_DMG_MOD             = 483
	SPA_FOCUS_INCOMING_DMG_AMT             = 484
	SPA_FOCUS_LIMIT_CASTER_CLASS           = 485
	SPA_FOCUS_LIMIT_SAME_CASTER            = 486
	SPA_EXTEND_TRADESKILL_CAP              = 487
	SPA_DEFENDER_MELEE_FORCE_PCT           = 488
	SPA_WORN_ENDURANCE_REGEN_CAP           = 489
	SPA_FOCUS_MIN_REUSE_TIME               = 490
	SPA_FOCUS_MAX_REUSE_TIME               = 491
	SPA_FOCUS_ENDURANCE_MIN                = 492
	SPA_FOCUS_ENDURANCE_MAX                = 493
	SPA_PET_ADD_ATK                        = 494
	SPA_FOCUS_DURATION_MAX                 = 495
	SPA_CRIT_MELEE_DMG_MOD_MAX             = 496
	SPA_FOCUS_CAST_PROC_NO_BYPASS          = 497
	SPA_ADD_EXTRA_PRIMARY_ATTACK_PCT       = 498
	SPA_ADD_EXTRA_SECONDARY_ATTACK_PCT     = 499
	SPA_FOCUS_CAST_TIME_MOD2               = 500
	SPA_FOCUS_CAST_TIME_AMT                = 501
	SPA_FEARSTUN                           = 502
	SPA_MELEE_DMG_POSITION_MOD             = 503
	SPA_MELEE_DMG_POSITION_AMT             = 504
	SPA_DMG_TAKEN_POSITION_MOD             = 505
	SPA_DMG_TAKEN_POSITION_AMT             = 506
	SPA_AMPLIFY_MOD                        = 507
	SPA_AMPLIFY_AMT                        = 508
	SPA_HEALTH_TRANSFER                    = 509
	SPA_FOCUS_RESIST_INCOMING              = 510
	SPA_FOCUS_TIMER_MIN                    = 511
	SPA_PROC_TIMER_MOD                     = 512
	SPA_MANA_MAX                           = 513
	SPA_ENDURANCE_MAX                      = 514
	SPA_AC_AVOIDANCE_MAX                   = 515
	SPA_AC_MITIGATION_MAX                  = 516
	SPA_ATTACK_OFFENSE_MAX                 = 517
	SPA_ATTACK_ACCURACY_MAX                = 518
	SPA_LUCK_AMT                           = 519
	SPA_LUCK_PCT                           = 520
	SPA_ENDURANCE_ABSORB_PCT_DMG           = 521
	SPA_INSTANT_MANA_PCT                   = 522
	SPA_INSTANT_ENDURANCE_PCT              = 523
	SPA_DURATION_HP_PCT                    = 524
	SPA_DURATION_MANA_PCT                  = 525
	SPA_DURATION_ENDURANCE_PCT             = 526

	SPA_CANCEL_MAGIC         = SPA_CANCEL_NEGATIVE_MAGIC
	SPA_NPC_REACTION_RATING  = SPA_NPC_AGGRO_RADIUS
	SPA_CLEAR_NPC_TARGETLIST = SPA_NPC_WIPE_HATE_LIST
)

const (
	SPELLCAT_AEGOLISM            = 1
	SPELLCAT_AGILITY             = 2
	SPELLCAT_ALLIANCE            = 3
	SPELLCAT_ANIMAL              = 4
	SPELLCAT_ANTONICA            = 5
	SPELLCAT_ARMOR_CLASS         = 6
	SPELLCAT_ATTACK              = 7
	SPELLCAT_BANE                = 8
	SPELLCAT_BLIND               = 9
	SPELLCAT_BLOCK               = 10
	SPELLCAT_CALM                = 11
	SPELLCAT_CHARISMA            = 12
	SPELLCAT_CHARM               = 13
	SPELLCAT_COLD                = 14
	SPELLCAT_COMBAT_ABILITIES    = 15
	SPELLCAT_COMBAT_INNATES      = 16
	SPELLCAT_CONVERSIONS         = 17
	SPELLCAT_CREATE_ITEM         = 18
	SPELLCAT_CURE                = 19
	SPELLCAT_DAMAGE_OVER_TIME    = 20
	SPELLCAT_DAMAGE_SHIELD       = 21
	SPELLCAT_DEFENSIVE           = 22
	SPELLCAT_DESTROY             = 23
	SPELLCAT_DEXTERITY           = 24
	SPELLCAT_DIRECT_DAMAGE       = 25
	SPELLCAT_DISARM_TRAPS        = 26
	SPELLCAT_DISCIPLINES         = 27
	SPELLCAT_DISCORD             = 28
	SPELLCAT_DISEASE             = 29
	SPELLCAT_DISEMPOWERING       = 30
	SPELLCAT_DISPEL              = 31
	SPELLCAT_DURATION_HEALS      = 32
	SPELLCAT_DURATION_TAP        = 33
	SPELLCAT_ENCHANT_METAL       = 34
	SPELLCAT_ENTHRALL            = 35
	SPELLCAT_FAYDWER             = 36
	SPELLCAT_FEAR                = 37
	SPELLCAT_FIRE                = 38
	SPELLCAT_FIZZLE_RATE         = 39
	SPELLCAT_FUMBLE              = 40
	SPELLCAT_HASTE               = 41
	SPELLCAT_HEALS               = 42
	SPELLCAT_HEALTH              = 43
	SPELLCAT_HEALTH_MANA         = 44
	SPELLCAT_HP_BUFFS            = 45
	SPELLCAT_HP_TYPE_ONE         = 46
	SPELLCAT_HP_TYPE_TWO         = 47
	SPELLCAT_ILLUSION_OTHER      = 48
	SPELLCAT_ILLUSION_ADVENTURER = 49
	SPELLCAT_IMBUE_GEM           = 50
	SPELLCAT_INVISIBILITY        = 51
	SPELLCAT_INVULNERABILITY     = 52
	SPELLCAT_JOLT                = 53
	SPELLCAT_KUNARK              = 54
	SPELLCAT_LEVITATE            = 55
	SPELLCAT_LIFE_FLOW           = 56
	SPELLCAT_LUCLIN              = 57
	SPELLCAT_MAGIC               = 58
	SPELLCAT_MANA                = 59
	SPELLCAT_MANA_DRAIN          = 60
	SPELLCAT_MANA_FLOW           = 61
	SPELLCAT_MELEE_GUARD         = 62
	SPELLCAT_MEMORY_BLUR         = 63
	SPELLCAT_MISC                = 64
	SPELLCAT_MOVEMENT            = 65
	SPELLCAT_OBJECTS             = 66
	SPELLCAT_ODUS                = 67
	SPELLCAT_OFFENSIVE           = 68
	SPELLCAT_PET                 = 69
	SPELLCAT_PET_HASTE           = 70
	SPELLCAT_PET_MISC_BUFFS      = 71
	SPELLCAT_PHYSICAL            = 72
	SPELLCAT_PICKLOCK            = 73
	SPELLCAT_PLANT               = 74
	SPELLCAT_POISON              = 75
	SPELLCAT_POWER_TAP           = 76
	SPELLCAT_QUICK_HEAL          = 77
	SPELLCAT_REFLECTION          = 78
	SPELLCAT_REGEN               = 79
	SPELLCAT_RESIST_BUFF         = 80
	SPELLCAT_RESIST_DEBUFFS      = 81
	SPELLCAT_RESURRECTION        = 82
	SPELLCAT_ROOT                = 83
	SPELLCAT_RUNE                = 84
	SPELLCAT_SENSE_TRAP          = 85
	SPELLCAT_SHADOWSTEP          = 86
	SPELLCAT_SHIELDING           = 87
	SPELLCAT_SLOW                = 88
	SPELLCAT_SNARE               = 89
	SPELLCAT_SPECIAL             = 90
	SPELLCAT_SPELL_FOCUS         = 91
	SPELLCAT_SPELL_GUARD         = 92
	SPELLCAT_SPELLSHIELD         = 93
	SPELLCAT_STAMINA             = 94
	SPELLCAT_STATISTIC_BUFFS     = 95
	SPELLCAT_STRENGTH            = 96
	SPELLCAT_STUN                = 97
	SPELLCAT_SUM_AIR             = 98
	SPELLCAT_SUM_ANIMATION       = 99
	SPELLCAT_SUM_EARTH           = 100
	SPELLCAT_SUM_FAMILIAR        = 101
	SPELLCAT_SUM_FIRE            = 102
	SPELLCAT_SUM_UNDEAD          = 103
	SPELLCAT_SUM_WARDER          = 104
	SPELLCAT_SUM_WATER           = 105
	SPELLCAT_SUMMON_ARMOR        = 106
	SPELLCAT_SUMMON_FOCUS        = 107
	SPELLCAT_SUMMON_FOOD_WATER   = 108
	SPELLCAT_SUMMON_UTILITY      = 109
	SPELLCAT_SUMMON_WEAPON       = 110
	SPELLCAT_SUMMONED            = 111
	SPELLCAT_SYMBOL              = 112
	SPELLCAT_TAELOSIA            = 113
	SPELLCAT_TAPS                = 114
	SPELLCAT_TECHNIQUES          = 115
	SPELLCAT_THE_PLANES          = 116
	SPELLCAT_TIMER_1             = 117
	SPELLCAT_TIMER_2             = 118
	SPELLCAT_TIMER_3             = 119
	SPELLCAT_TIMER_4             = 120
	SPELLCAT_TIMER_5             = 121
	SPELLCAT_TIMER_6             = 122
	SPELLCAT_TRANSPORT           = 123
	SPELLCAT_UNDEAD              = 124
	SPELLCAT_UTILITY_BENEFICIAL  = 125
	SPELLCAT_UTILITY_DETRIMENTAL = 126
	SPELLCAT_VELIOUS             = 127
	SPELLCAT_VISAGES             = 128
	SPELLCAT_VISION              = 129
	SPELLCAT_WISDOM_INTELLIGENCE = 130
	SPELLCAT_TRAPS               = 131
	SPELLCAT_AURAS               = 132
	SPELLCAT_ENDURANCE           = 133
	SPELLCAT_SERPENTS_SPINE      = 134
	SPELLCAT_CORRUPTION          = 135
	SPELLCAT_LEARNING            = 136
	SPELLCAT_CHROMATIC           = 137
	SPELLCAT_PRISMATIC           = 138
	SPELLCAT_SUM_SWARM           = 139
	SPELLCAT_DELAYED             = 140
	SPELLCAT_TEMPORARY           = 141
	SPELLCAT_TWINCAST            = 142
	SPELLCAT_SUM_BODYGUARD       = 143
	SPELLCAT_HUMANOID            = 144
	SPELLCAT_HASTE_SPELL_FOCUS   = 145
	SPELLCAT_TIMER_7             = 146
	SPELLCAT_TIMER_8             = 147
	SPELLCAT_TIMER_9             = 148
	SPELLCAT_TIMER_10            = 149
	SPELLCAT_TIMER_11            = 150
	SPELLCAT_TIMER_12            = 151
	SPELLCAT_HATRED              = 152
	SPELLCAT_FAST                = 153
	SPELLCAT_ILLUSION_SPECIAL    = 154
	SPELLCAT_TIMER_13            = 155
	SPELLCAT_TIMER_14            = 156
	SPELLCAT_TIMER_15            = 157
	SPELLCAT_TIMER_16            = 158
	SPELLCAT_TIMER_17            = 159
	SPELLCAT_TIMER_18            = 160
	SPELLCAT_TIMER_19            = 161
	SPELLCAT_TIMER_20            = 162
	SPELLCAT_ALARIS              = 163
	SPELLCAT_COMBINATION         = 164
	SPELLCAT_INDEPENDENT         = 165
	SPELLCAT_SKILL_ATTACKS       = 166
	SPELLCAT_INCOMING            = 167
	SPELLCAT_CURSE               = 168
	SPELLCAT_TIMER_21            = 169
	SPELLCAT_TIMER_22            = 170
	SPELLCAT_TIMER_23            = 171
	SPELLCAT_TIMER_24            = 172
	SPELLCAT_TIMER_25            = 173
	SPELLCAT_DRUNKENNESS         = 174
	SPELLCAT_THROWING            = 175
	SPELLCAT_MELEE_DAMAGE        = 176
)

func IsSpellCountersSPA(attrib int) bool {
	return attrib == SPA_DISEASE || attrib == SPA_POISON || attrib == SPA_CURSE || attrib == SPA_CORRUPTION
}

func IsDamageAbsorbSPA(attrib int) bool {
	return attrib == SPA_STONESKIN || attrib == SPA_SPELL_SHIELD || attrib == SPA_SPELL_GUARD || attrib == SPA_MELEE_GUARD // Mitigate Melee Damage || attrib == SPA_DOT_GUARD                   // DoT Guard || attrib == SPA_MELEE_THRESHOLD_GUARD       // Melee Threshold Guard || attrib == SPA_SPELL_THRESHOLD_GUARD;      // Spell Threshold Guard
}

const // eResistType
(
	ResistType_None       = 0
	ResistType_Magic      = 1
	ResistType_Fire       = 2
	ResistType_Cold       = 3
	ResistType_Poison     = 4
	ResistType_Disease    = 5
	ResistType_Chromatic  = 6
	ResistType_Prismatic  = 7
	ResistType_Physical   = 8
	ResistType_Corruption = 9
)

const (
	SpellType_Detrimental         = 0
	SpellType_Beneficial          = 1
	SpellType_BeneficialGroupOnly = 2
)

// Determines the algorithm used to affect the spell value potentially affected by
// time or by level or other things too...
const // eSpellValueRangeCalc
(
	SpellValueRangeCalc_DecayTick1  = 107
	SpellValueRangeCalc_DecayTick2  = 108
	SpellValueRangeCalc_DecayTick5  = 120
	SpellValueRangeCalc_DecayTick12 = 122
	SpellValueRangeCalc_Random      = 123
)

// const // eSpellNoOverwrite : int
// (
// 	NoOverwrite_Default     // Spell can be overwritten normally
// 	NoOverwrite_OtherSpells // Spell cannot be overwritten by other spells except itself
// 	NoOverwrite_AllSpells   // Spell cannot be overwritten by any spell
// )

// const // eSpellRecourseType : int
// (
// 	SpellRecourseType_AlwaysHit         // Recourse for every target that it could hit
// 	SpellRecourseType_AlwaysHitNoResist // Recourse for every target
// 	SpellRecourseType_OnceNoResist      // Recourse once if target exists
// 	SpellRecourseType_Once              // Recourse once if spell hits
// )

const // eSpellTargetType : uint8_t
(
	TargetType_None             = 0
	TargetType_LineOfSight      = 1
	TargetType_AEPC_v1          = 2 // players in area around caster
	TargetType_Group_v1         = 3 // group members around caster
	TargetType_PBAE             = 4 // area around caster
	TargetType_Single           = 5 // current target
	TargetType_Self             = 6 // targets self only
	TargetType_TargetArea       = 8 // radius around target
	TargetType_TargetAnimal     = 9
	TargetType_TargetUndead     = 10
	TargetType_TargetSummoned   = 11
	TargetType_TargetDrain      = 13
	TargetType_Pet              = 14 // caster's pet
	TargetType_TargetCorpse     = 15
	TargetType_TargetPlant      = 16
	TargetType_TargetGiants     = 17
	TargetType_TargetDragons    = 18
	TargetType_TargetColdain    = 19
	TargetType_TargetAEDrain    = 20
	TargetType_TargetAEUndead   = 24
	TargetType_TargetAESummoned = 25
	TargetType_HateList         = 32 // all players on hatelist in range
	TargetType_HateList_All     = 33 // all players on hatelist regardless of range
	TargetType_TargetCursed     = 34
	TargetType_TargetMuramite   = 35
	TargetType_CasterAreaPC     = 36
	TargetType_CasterAreaNPC    = 37
	TargetType_Pet_v2           = 38 // targeted pet
	TargetType_TargetPC         = 39 // targeted player
	TargetType_AEPC_v2          = 40 // area beneficial players
	TargetType_Group_v2         = 41 // area grouped players
	TargetType_DirectionalCone  = 42 // projected cone in front of player
	TargetType_SingleGrouped    = 43 // single target grouped
	TargetType_Beam             = 44
	TargetType_FreeTarget       = 45 // player picks a point in space
	TargetType_TargetOfTarget   = 46
	TargetType_PetOwner         = 47 // cast on pet's owner
	TargetType_AreaDetrimental  = 50 // targets enemies of caster
	TargetType_TargetBeneficial = 52
)

// const // eSpellStringType
// (
// 	SpellStringCastByMe
// 	SpellStringCastByOther
// 	SpellStringCastOnYou
// 	SpellStringCastOnAnother
// 	SpellStringWearOff
// )

const (
	MAX_SPELL_REAGENTS = 4
)

// const // EEffectActor
// (
// 	EEA_None
// 	EEA_Caster
// 	EEA_Missile
// 	EEA_Target
// 	EEA_COUNT
// )

// const // EAttachPoint
// (
// 	EAP_None
// 	EAP_Default
// 	EAP_Chest
// 	EAP_Head
// 	EAP_LeftHand
// 	EAP_RightHand
// 	EAP_LeftFoot
// 	EAP_RightFoot
// 	EAP_Weapon
// 	EAP_LeftEye
// 	EAP_RightEye
// 	EAP_Mouth
// 	EAP_Ground
// 	EAP_Cnt
// )

// //Matching stack group ID rules
// const // ESpellStackingRules
// (
// 	ESSR_None // default
// 	ESSR_SingleCaster
// 	ESSR_AllCasters
// 	ESSR_SingleCasterOnlyGreater
// 	ESSR_AllCastersOnlyGreater
// 	ESSR_SingleCasterNeverOverwrite
// 	ESSR_AllCastersNeverOverwrite
// 	ESSR_SingleCasterAlwaysOverwrite
// 	ESSR_AllCastersAlwaysOverwrite
// 	ESSR_Invalid
// )

const (
	MAX_SPELLEFFECTS = 999
)

var szSPATypes = []string{

	0:   "HP",
	1:   "AC",
	2:   "ATK",
	3:   "Movement Rate",
	4:   "STR",
	5:   "DEX",
	6:   "AGI",
	7:   "STA",
	8:   "INT",
	9:   "WIS",
	10:  "CHA",
	11:  "Melee Speed",
	12:  "Invisibility",
	13:  "See Invisible",
	14:  "Water Breathing",
	15:  "Mana",
	16:  "NPC Frenzy",
	17:  "NPC Awareness",
	18:  "Pacify",
	19:  "NPC Faction",
	20:  "Blindness",
	21:  "Stun",
	22:  "Charm",
	23:  "Fear",
	24:  "Stamina Loss",
	25:  "Bind Affinity",
	26:  "Gate",
	27:  "Dispel Magic",
	28:  "Invisibility Vs Undead",
	29:  "Invisibility Vs Animals",
	30:  "NPC Aggro Radius",
	31:  "Mesmerize",
	32:  "Summon",
	33:  "Summon Pet",
	34:  "Confuse",
	35:  "Disease Counter",
	36:  "Poison Counter",
	37:  "Detect Hostile",
	38:  "Detect Magic",
	39:  "Stacking: No Twincast",
	40:  "Invulnerability",
	41:  "Banish",
	42:  "Shadow Step",
	43:  "Berserk",
	44:  "Lycanthropy",
	45:  "Vampirism",
	46:  "Fire Resist",
	47:  "Cold Resist",
	48:  "Poison Resist",
	49:  "Disease Resist",
	50:  "Magic Resist",
	51:  "Detect Traps",
	52:  "Detect Undead",
	53:  "Detect Summoned",
	54:  "Detect Animals",
	55:  "Absorb Damage",
	56:  "True North",
	57:  "Levitation",
	58:  "Illusion",
	59:  "Damage Shield",
	60:  "Transfer Item",
	61:  "Identify",
	62:  "Item ID",
	63:  "Memblur",
	64:  "Spin Stun",
	65:  "Infravision",
	66:  "Ultravision",
	67:  "Eye Of Zomm",
	68:  "Reclaim Energy",
	69:  "Max HP",
	70:  "Corpse Bomb",
	71:  "Create Undead Pet",
	72:  "Preserve Corpse",
	73:  "Bind Sight",
	74:  "Feign Death",
	75:  "Ventriloquism",
	76:  "Sentinel",
	77:  "Locate Corpse",
	78:  "SpellShield",
	79:  "HP",
	80:  "Enchant:Light",
	81:  "Resurrect",
	82:  "Summon Player",
	83:  "Teleport",
	84:  "Toss",
	85:  "Add Proc",
	86:  "Reaction Radius",
	87:  "Magnification",
	88:  "Evacuate",
	89:  "Player Size",
	90:  "Ignore Pet",
	91:  "Summon Corpse",
	92:  "Hate",
	93:  "Control Weather",
	94:  "Make Fragile",
	95:  "Sacrifice",
	96:  "Silence",
	97:  "Max Mana",
	98:  "Bard Haste",
	99:  "Root",
	100: "HoT Heals",
	101: "Complete Heal (with duration)",
	102: "Pet Fearless",
	103: "Summon Pet",
	104: "Translocate",
	105: "Anti Gate",
	106: "Summon Warder",
	107: "Alter NPC Level",
	108: "Summon Familiar",
	109: "Summon In Bag",
	110: "Archery",
	111: "All Resists",
	112: "Casting Level",
	113: "Summon Mount",
	114: "Hate Multiplier",
	115: "Food/Water",
	116: "Curse Counter",
	117: "Make Weapons Magical",
	118: "Singing Skill",
	119: "Melee Overhaste",
	120: "Healing Effectiveness",
	121: "Reverse Damage Shield",
	122: "Reduce Skill",
	123: "Immunity",
	124: "Spell Damage",
	125: "Healing",
	126: "Spell Resist Rate",
	127: "Spell Cast Time",
	128: "Spell Duration",
	129: "Spell Range",
	130: "Spell/Bash Hate",
	131: "Chance of Using Reagent",
	132: "Spell Mana Cost",
	133: "Spell Stun Duration",
	134: "Limit: Max Level",
	135: "Limit: Resist",
	136: "Limit: Target",
	137: "Limit: Effect",
	138: "Limit: SpellType",
	139: "Limit: Spell",
	140: "Limit: Min Duration",
	141: "Limit: Instant spells only",
	142: "Limit: Min Level",
	143: "Limit: Min Cast Time",
	144: "Limit: Max Cast Time",
	145: "NPC Warder Banish",
	146: "Portal Locations",
	147: "HP",
	148: "Stacking: Block",
	149: "Stacking: Overwrite",
	150: "Death Save",
	151: "Pocket Pet",
	152: "Summon a Pet Swarm",
	153: "Balance Party Damage",
	154: "Remove Detrimental",
	155: "PoP Resurrect",
	156: "Mirror",
	157: "Spell Damage Shield",
	158: "Reflect Spell",
	159: "All Stats",
	160: "Drunk",
	161: "Mitigate Spell Damage",
	162: "Mitigate Melee Damage",
	163: "Absorb Damage",
	164: "Sense LDoN Chest",
	165: "Disarm LDoN Trap",
	166: "Unlock LDoN Chest",
	167: "Increase Pet Power",
	168: "Defensive",
	169: "Chance to Critical Melee",
	170: "Critical Direct Damage",
	171: "Chance to Crippling Blow",
	172: "Evasion",
	173: "Riposte",
	174: "Dodge",
	175: "Parry",
	176: "Dual Wield",
	177: "Double Attack",
	178: "Melee Lifetap",
	179: "Instrument Modifier",
	180: "Chance to Resist Spells",
	181: "Chance to Resist Fear Spell",
	182: "Melee Attack",
	183: "Skill Chance",
	184: "Chance to Hit",
	185: "Skills Damage Modifier",
	186: "Skills Minimum Damage Modifier",
	187: "Balance Party Mana",
	188: "Chance to block",
	189: "Endurance",
	190: "Max Endurance",
	191: "Amnesia",
	192: "Hate",
	193: "Skill Attack",
	194: "Fade",
	195: "Stun Resist",
	196: "Strikethrough",
	197: "Skill Damage Taken",
	198: "Instant Endurance",
	199: "Taunt",
	200: "Proc Chance",
	201: "Ranged Proc",
	202: "Illusion Other",
	203: "Mass Group Buff",
	204: "Group Fear Immunity",
	205: "AE Rampage",
	206: "AE Taunt",
	207: "Flesh to Bone",
	208: "Purge Poison",
	209: "Cancel Beneficial",
	210: "Pet Shield",
	211: "AE Melee",
	212: "Frenzied Devastation",
	213: "Pet HP",
	214: "Change Max HP",
	215: "Pet Avoidance",
	216: "Accuracy",
	217: "Headshot",
	218: "Pet Crit Melee",
	219: "Slay Undead",
	220: "Damage Bonus",
	221: "Reduce Weight",
	222: "Block Behind",
	223: "Double Riposte",
	224: "Additional Riposte",
	225: "Double Attack",
	226: "2H bash",
	227: "Base Refresh Timer",
	228: "Reduce Fall Dmg",
	229: "Cast Through Stun",
	230: "Increase Shield Dist",
	231: "Stun Bash Chance",
	232: "Divine Save",
	233: "Metabolism",
	234: "Poison Mastery",
	235: "Focus Channelling",
	236: "Free Pet",
	237: "Pet Affinity",
	238: "Permanent Illusion",
	239: "Stonewall",
	240: "String Unbreakable",
	241: "Improve Reclaim Energy",
	242: "Increase Chance Memwipe",
	243: "NoBreak Charm Chance",
	244: "Root Break Chance",
	245: "Trap Circumvention",
	246: "Lung Capacity",
	247: "Increase SkillCap",
	248: "Extra Specialization",
	249: "Offhand Min",
	250: "Spell Proc Chance",
	251: "Endless Quiver",
	252: "Backstab from Front",
	253: "Chaotic Stab",
	254: "NoSpell",
	255: "Shielding Duration",
	256: "Shroud Of Stealth",
	257: "Give Pet Hold",
	258: "Triple Backstab",
	259: "AC Limit",
	260: "Add Instrument",
	261: "Song Cap",
	262: "Cap",
	263: "Tradeskill Masteries",
	264: "Reduce AATimer",
	265: "No Fizzle",
	266: "Add Attack Chance",
	267: "Add Pet Commands",
	268: "Failure Rate",
	269: "Bandage Perc Limit",
	270: "Bard Song Range",
	271: "Base Run Speed",
	272: "Casting Level",
	273: "Critical DoT Chance",
	274: "Critical Heal Chance",
	275: "Critical Mend",
	276: "Dual Wield Amt",
	277: "Extra DI Chance",
	278: "Finishing Blow",
	279: "Flurry Chance",
	280: "Pet Flurry Chance",
	281: "Give Pet Feign",
	282: "Bandage Amount",
	283: "Special Attack Chain",
	284: "LoH Set Heal",
	285: "NoMove Check Sneak",
	286: "DD Bonus",
	287: "Focus Combat Duration",
	288: "Add Proc Hit",
	289: "Trigger Effect",
	290: "Movement Cap",
	291: "Purify",
	292: "Strikethrough2",
	293: "StunResist2",
	294: "Critical DD Chance",
	295: "Reduce Timer Special",
	296: "Incoming Spell Damage",
	297: "Incoming Spell Damage Amt",
	298: "Pet Height",
	299: "Wake the Dead",
	300: "Doppelganger",
	301: "Range Damage",
	302: "Damage Crit",
	303: "Damage",
	304: "Secondary Riposte",
	305: "Damage Shield Mitigation",
	306: "Wake The Dead 2",
	307: "Appraisal",
	308: "Zone Suspend Minion",
	309: "Gate Caster's Bindpoint",
	310: "Decrease Reuse Timer",
	311: "Limit: Combat Skills Not Allowed",
	312: "Observer",
	313: "Forage Master",
	314: "Improved Invisibility",
	315: "Improved Invisibility Vs Undead",
	316: "Improved Invisibility Vs Animals",
	317: "Worn Regen Cap",
	318: "Worn Mana Cap",
	319: "Critical HP Regen",
	320: "Shield Block Chance",
	321: "Reduce Target Hate",
	322: "Gate to Starting City",
	323: "Add Defensive Proc",
	324: "HP for Mana",
	325: "No Break AE Sneak",
	326: "Spell Slots",
	327: "Buff Slots",
	328: "Negative HP Limit",
	329: "Mana Shield Absorb Damage",
	330: "Critical Damage",
	331: "Item Recovery",
	332: "Summon to Corpse",
	333: "Trigger Effect",
	334: "HP",
	335: "Block Next Spell",
	336: "Illusionary Target",
	337: "Experience",
	338: "Expedient Recovery",
	339: "Trigger DoT",
	340: "Trigger DD",
	341: "Worn Attack Cap",
	342: "Prevent Flee on Low Health",
	343: "Spell Interrupt",
	344: "Item Channeling",
	345: "Assassinate Max",
	346: "Headshot Max",
	347: "Double Ranged Attack",
	348: "Limit: Min Mana",
	349: "Damage With Shield",
	350: "Manaburn",
	351: "Persistent Effect",
	352: "Trap Count",
	353: "SOI Count",
	354: "Deactivate All Traps",
	355: "Learn Trap",
	356: "Change Trigger Type",
	357: "Mute",
	358: "Mana/Max Mana",
	359: "Passive Sense Trap",
	360: "Proc On Kill Shot",
	361: "Proc On Death",
	362: "Potion Belt",
	363: "Bandolier",
	364: "Triple Attack Chance",
	365: "Trigger on Kill Shot",
	366: "Group Shielding",
	367: "Modify Body Type",
	368: "Modify Faction",
	369: "Corruption Counter",
	370: "Corruption Resist",
	371: "Slow",
	372: "Grant Foraging",
	373: "Trigger Effect",
	374: "Trigger Spell",
	375: "Critical DoT Damage",
	376: "Fling",
	377: "Trigger Effect",
	378: "Resist",
	379: "Directional Shadowstep",
	380: "Knockback",
	381: "Fling to Self",
	382: "Negate",
	383: "Trigger Spell",
	384: "Fling to Target",
	385: "Limit: SpellGroup",
	386: "Trigger Effect",
	387: "Trigger Effect",
	388: "Summon All Corpses",
	389: "Spell Gem Refresh",
	390: "Spell Gem Lockout",
	391: "Limit: Max Mana",
	392: "Instant Heal Amt",
	393: "Incoming Healing Effectiveness",
	394: "Incoming Healing Amt",
	395: "Heal Crit",
	396: "Heal Crit Amt",
	397: "Pet Amt Mitigation",
	398: "Swarm Pet Duration",
	399: "Twincast Chance",
	400: "Healburn",
	401: "Mana/HP",
	402: "Endurance/HP",
	403: "Limit: SpellClass",
	404: "Limit: SpellSubclass",
	405: "Staff Block Chance",
	406: "Trigger Effect",
	407: "Trigger Effect",
	408: "HP Limit",
	409: "Mana Limit",
	410: "Endurance Limit",
	411: "Limit: PlayerClass",
	412: "Limit: Race",
	413: "Base Damage",
	414: "Limit: CastingSkill",
	415: "Limit: ItemClass",
	416: "AC",
	417: "Mana Regen",
	418: "Skill Damage",
	419: "Add Proc",
	420: "Limit: Use",
	421: "Limit: Use Amt",
	422: "Limit: Use Min",
	423: "Limit: Use Type",
	424: "Gravitate",
	425: "Fly",
	426: "AddExtTargetSlots",
	427: "Skill Proc",
	428: "Limit Skill",
	429: "Hit Limited Proc",
	430: "PostEffect",
	431: "PostEffectData",
	432: "ExpandMaxActiveTrophyBenefits",
	433: "Min Damage",
	434: "Min Damage",
	435: "Fragile Defense",
	436: "Beneficial Countdown Hold",
	437: "Teleport",
	438: "Translocate",
	439: "Assassinate",
	440: "FinishingBlowMax",
	441: "Distance Removal",
	442: "Trigger Effect",
	443: "Trigger Effect",
	444: "Improved Taunt",
	445: "Add Merc Slot",
	446: "A_Stacker",
	447: "B_Stacker",
	448: "C_Stacker",
	449: "D_Stacker",
	450: "DoT Guard",
	451: "Melee Threshold Guard",
	452: "Spell Threshold Guard",
	453: "Trigger Effect",
	454: "Trigger Effect",
	455: "Add Hate",
	456: "Add Hate Over Time",
	457: "Resource Tap",
	458: "Faction",
	459: "Skill Damage Mod 2",
	460: "Limit: Include Non-Focusable",
	461: "Spell Damage 2",
	462: "Spell Damage Amt 2",
	463: "Shield Target",
	464: "PC Pet Rampage",
	465: "PC Pet AE Rampage",
	466: "PC Pet Flurry Chance",
	467: "DS Mitigation Amt",
	468: "DS Mitigation Percentage",
	469: "Chance Best in Spell Group",
	470: "Trigger Best in Spell Group",
	471: "Double Melee Round(PC Only)",
	472: "Toggle Passive AA Rank",
	473: "Double Backstab From Front",
	474: "Pet Crit Melee Damage",
	475: "Trigger Spell Non-Item",
	476: "Weapon Stance",
	477: "Move to Top of Hatelist",
	478: "Move to Bottom of Hatelist",
	479: "Limit: Value",
	480: "Limit: Value",
	481: "Trigger Spell",
	482: "Base Hit Damage",
	483: "Spell Damage taken",
	484: "Spell Damage taken",
	485: "Limit: CasterClass",
	486: "Limit: Caster",
	487: "Extend Tradeskill Cap",
	488: "Push Taken",
	489: "Worn Endurance Regen Cap",
	490: "Limit: ReuseTime Min",
	491: "Limit: ReuseTime Max",
	492: "Limit: Endurance Min",
	493: "Limit: Endurance Max",
	494: "Pet Add Attack",
	495: "Limit: Duration Max",
	496: "Critical Hit Damage",
	497: "NoProc",
	498: "Extra Attack % (1H Primary)",
	499: "Extra Attack % (1H Secondary)",
	500: "Spell Haste v2",
	501: "Spell Cast Time",
	502: "Stun and Fear",
	503: "Rear Arc Melee Damage Mod",
	504: "Rear Arc Melee Damage",
	505: "Rear Arc Damage Taken Mod",
	506: "Rear Arc Damage Taken",
	507: "Spell Damage v4 Mod",
	508: "Spell Damage v4",
	509: "Health Transfer",
	510: "Resist Incoming",
	511: "Focus Timer Min",
	512: "Proc Timer Modifier",
	513: "Mana Max Percent",
	514: "Endurance Max Percent",
	515: "AC Avoidance Max Percent",
	516: "AC Mitigation Max Percent",
	517: "Attack Offense Max Percent",
	518: "Attack Accuracy Max Percent",
	519: "Luck Amt",
	520: "Luck Percent",
}

const (
	Warrior      = 1
	Cleric       = 2
	Paladin      = 3
	Ranger       = 4
	Shadowknight = 5
	Druid        = 6
	Monk         = 7
	Bard         = 8
	Rogue        = 9
	Shaman       = 10
	Necromancer  = 11
	Wizard       = 12
	Mage         = 13
	Enchanter    = 14
	Beastlord    = 15
	Berserker    = 16
)

var classNames = []string{
	1:  "Warrior",
	2:  "Cleric",
	3:  "Paladin",
	4:  "Ranger",
	5:  "Shadowknight",
	6:  "Druid",
	7:  "Monk",
	8:  "Bard",
	9:  "Rogue",
	10: "Shaman",
	11: "Necromancer",
	12: "Wizard",
	13: "Mage",
	14: "Enchanter",
	15: "Beastlord",
	16: "Berserker",
}

const (
	TT_PBAE        = 4
	TT_TARGETED_AE = 8
	TT_AE_PC_V2    = 40
	TT_DIRECTIONAL = 42
	TT_SPLASH      = 45
)

var spellRestrictions = []string{
	0:     "None",
	100:   "Only works on Animal or Humanoid",
	101:   "Only works on Dragon",
	102:   "Only works on Animal or Insect",
	104:   "Only works on Animal",
	105:   "Only works on Plant",
	106:   "Only works on Giant",
	108:   "Doesn't work on Animals or Humanoids",
	109:   "Only works on Bixie",
	110:   "Only works on Harpy",
	111:   "Only works on Gnoll",
	112:   "Only works on Sporali",
	113:   "Only works on Kobold",
	114:   "Only works on Shade",
	115:   "Only works on Drakkin",
	117:   "Only works on Animals or Plants",
	118:   "Only works on Summoned",
	119:   "Only works on Fire_Pet",
	120:   "Only works on Undead",
	121:   "Only works on Living",
	122:   "Only works on Fairy",
	123:   "Only works on Humanoid",
	124:   "Undead HP Less Than 10%",
	125:   "Clockwork HP Less Than 45%",
	126:   "Wisp HP Less Than 10%",
	190:   "Doesn't work on Raid Bosses",
	191:   "Only works on Raid Bosses",
	201:   "HP Above 75%",
	203:   "HP Less Than 20%",
	204:   "HP Less Than 50%",
	216:   "Not In Combat",
	221:   "At Least 1 Pet On Hatelist",
	222:   "At Least 2 Pets On Hatelist",
	223:   "At Least 3 Pets On Hatelist",
	224:   "At Least 4 Pets On Hatelist",
	225:   "At Least 5 Pets On Hatelist",
	226:   "At Least 6 Pets On Hatelist",
	227:   "At Least 7 Pets On Hatelist",
	228:   "At Least 8 Pets On Hatelist",
	229:   "At Least 9 Pets On Hatelist",
	230:   "At Least 10 Pets On Hatelist",
	231:   "At Least 11 Pets On Hatelist",
	232:   "At Least 12 Pets On Hatelist",
	233:   "At Least 13 Pets On Hatelist",
	234:   "At Least 14 Pets On Hatelist",
	235:   "At Least 15 Pets On Hatelist",
	236:   "At Least 16 Pets On Hatelist",
	237:   "At Least 17 Pets On Hatelist",
	238:   "At Least 18 Pets On Hatelist",
	239:   "At Least 19 Pets On Hatelist",
	240:   "At Least 20 Pets On Hatelist",
	250:   "HP Less Than 35%",
	304:   "Chain Plate Classes",
	399:   "HP Between 15 and 25%",
	400:   "HP Between 1 and 25%",
	401:   "HP Between 25 and 35%",
	402:   "HP Between 35 and 45%",
	403:   "HP Between 45 and 55%",
	404:   "HP Between 55 and 65%",
	412:   "HP Above 99%",
	501:   "HP Below 5%",
	502:   "HP Below 10%",
	503:   "HP Below 15%",
	504:   "HP Below 20%",
	505:   "HP Below 25%",
	506:   "HP Below 30%",
	507:   "HP Below 35%",
	508:   "HP Below 40%",
	509:   "HP Below 45%",
	510:   "HP Below 50%",
	511:   "HP Below 55%",
	512:   "HP Below 60%",
	513:   "HP Below 65%",
	514:   "HP Below 70%",
	515:   "HP Below 75%",
	516:   "HP Below 80%",
	517:   "HP Below 85%",
	518:   "HP Below 90%",
	519:   "HP Below 95%",
	521:   "Mana Below X%",
	522:   "End Below 40%",
	523:   "Mana Below 40%",
	603:   "Only works on Undead2",
	608:   "Only works on Undead3",
	624:   "Only works on Summoned2",
	701:   "Doesn't work on Pets",
	818:   "Only works on Undead4",
	819:   "Doesn't work on Undead4",
	825:   "End Below 21%",
	826:   "End Below 25%",
	827:   "End Below 29%",
	836:   "Only works on Regular Servers",
	837:   "Doesn't work on Progression Servers",
	842:   "Only works on Humanoid Level 84 Max",
	843:   "Only works on Humanoid Level 86 Max",
	844:   "Only works on Humanoid Level 88 Max",
	1000:  "Between Level 1 and 75",
	1001:  "Between Level 76 and 85",
	1002:  "Between Level 86 and 95",
	1003:  "Between Level 96 and 100",
	1004:  "HP Less Than 80%",
	38311: "Mana Below 20%",
	38312: "Mana Below 10%",
}
