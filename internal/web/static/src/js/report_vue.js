if ($(".report").length) {
     let report = new Vue({
        el: '.report__details',
        data: {
            total_messages: data.report_analytics.MessagesSent,
            total_words: data.report_analytics.WordsSent,
            hour_map: data.report_analytics.HourList,
            participants: data.report_analytics.Participants,
        },

        computed: {
            avg_messages_per_day: function() {
                let maxMonths = 0;
                for (var participant in this.participants) {
                    if (this.participants[participant].MonthList.length > maxMonths) {
                        maxMonths = this.participants[participant].MonthList.length
                    }
                }
                return Math.floor(this.total_messages / maxMonths / 30)
            },
            active_months: function() {
                let maxMonths = 0;
                for (var participant in this.participants) {
                    if (this.participants[participant].MonthList.length > maxMonths) {
                        maxMonths = this.participants[participant].MonthList.length
                    }
                }
                return maxMonths
            },
            busiest_hour: function () {
                let max = this.hour_map.reduce((prev, current) => {
                    return (prev.Messages > current.Messages) ? prev : current
                })
                return max
            },
            messages_pie: function () {
                let pie = [];
                for (var p in this.participants) {
                    pie.push([p, this.participants[p].MessagesSent])
                }
                return pie;
            },
            word_pie: function () {
                let pie = [];
                for (var p in this.participants) {
                    pie.push([p, this.participants[p].WordsSent])
                }
                return pie
            },
            months: function () {
                let participant_with_most_months = {}
                for (let participant in this.participants) {
                    if (typeof participant_with_most_months.MonthList === "undefined") {
                        participant_with_most_months = this.participants[participant]
                        continue
                    }
                    if (this.participants[participant].MonthList.length > participant_with_most_months.MonthList.length) {
                        participant_with_most_months = this.participants[participant]
                    }
                }
                return Object.keys(participant_with_most_months.MonthList).map(function(key) {
                    return participant_with_most_months.MonthList[key].Month;
                });
            },
            messages_per_month: function () {
                let series = [];
                for (let p in this.participants){
                    let data = {}
                    data.name = p
                    data.data = [];
                    series.push(data)

                    for (let month in this.participants[p].MonthList) {
                        data.data.push([this.months.indexOf(this.participants[p].MonthList[month].Month), this.participants[p].MonthList[month].Messages])
                    }
                }
                return series
            }
            // compiledMarkdown: function () {
            //     return marked(this.input, {sanitize: true})
            // }
        },
        methods: {
            get_top_emoji: function (wordlist) {
                for (var word in wordlist) {
                    if (is_emoji(wordlist[word].Word)) {
                        console.log("Found most used emoji:" + wordlist[word].Word)
                        return wordlist[word].Word
                    }
                }
            }
        },
        mounted() {
            MessagePie(this.messages_pie)
            WordPie(this.word_pie)
            MonthlyMessages(this.months, this.messages_per_month)
        }
    })



}
