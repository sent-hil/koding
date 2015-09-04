module.exports = {
  getters : require './getters'

  actions :
    emoji   : require './actions/emoji'
    channel : require './actions/channel'

  stores  : [
    require './stores/emoji/emojisstore'
    require './stores/emoji/filteredemojilistquerystore'
    require './stores/emoji/filteredemojilistselectedindexstore'
    require './stores/emoji/commonemojilistselectedindexstore'
    require './stores/emoji/commonemojilistvisibilitystore'
    require './stores/channel/chatinputchannelsquerystore'
    require './stores/channel/chatinputchannelsselectedindexstore'
    require './stores/channel/chatinputchannelsvisibilitystore'
  ]
}

