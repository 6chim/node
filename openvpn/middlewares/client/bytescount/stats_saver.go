/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package bytescount

import "github.com/mysteriumnetwork/node/client/stats"

// NewSessionStatsSaver returns stats handler, which saves stats stats keeper
func NewSessionStatsSaver(statsKeeper stats.SessionStatsKeeper) SessionStatsHandler {
	return func(sessionStats stats.SessionStats) error {
		statsKeeper.Save(sessionStats)
		return nil
	}
}
