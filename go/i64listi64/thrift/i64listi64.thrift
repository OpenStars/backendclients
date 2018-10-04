namespace java OpenStars.Common.list.i64i64
namespace cpp OpenStars.Common.List.I64I64
namespace go OpenStars.Common.list.i64listi64

typedef i64 TEntry
typedef i64 TKey

typedef list<TKey> TKeyList
typedef list<TEntry> TEntryList
struct TValue {
	1: TEntryList entries;
}

enum LimitPolicy {
	REJECT = 0,
	EVICT_FRONT,
	EVICT_BACK
}

service I64ListI64Service {
	/**
	 * Count entry number
	 */
	i32			count(1: TKey key);

	/**
	 * Test entry existence
	 */
	bool		existed(1: TKey key, 2: TEntry entry);

	/**
	 * Get list of all entries
	 */
	TEntryList	getListAll(1: TKey key);
	TEntryList	getAll(1: TKey key);

	/**
	 * Get list entry slice
	 */
	TEntryList	getSlice(1: TKey key, 2: i32 start, 3: i32 length);
	TEntryList	getSliceFromEntry(1: TKey key, 2: TEntry entryStart, 3: i32 length);
	TEntryList	getSliceFromId(1: TKey key, 2: TEntry entryStart, 3: i32 length);

	//Get list entry slice
	//start: zero-based index count from back
	TEntryList	getSliceReverse(1: TKey key, 2: i32 start, 3: i32 length);
	TEntryList	getSliceFromEntryReverse(1: TKey key, 2: TEntry entryStart, 3: i32 length);
	TEntryList	getSliceFromIdReverse(1: TKey key, 2: TEntry entryStart, 3: i32 length);

	/**
	 * Put an entry, update on entry existence
	 * For unsorted list this insert at end()
	 */
	bool		put(1: TKey key, 2: TEntry entry);

	/**
	 * Clear old data, put new value on the key
	 */
	bool		putValue(1: TKey key, 2: TValue value);

	/**
	 * Put list of entry, update on entry existence
	 */
	bool		putList(1: TKey key, 2: TEntryList entryList);

	/**
	 * Insert an entry, return false on key existence
	 * @return true on successful insertion, false otherwise (due to length limit or entry existed on sorted list)
	 * For unsorted list this insert at end()
	 */
	bool		insert(1: TKey key, 2: TEntry entry);

	/**
	 * Insert an entry at index
	 * @param index considered count from begin if non-negative, count from end otherwise
	 * the entry will have index = start after insertion
	 * e.g. 0 = push_front, -1 = push_back, -2 = the entry before last entry
	 * For sorted list this calls insert
	 */
	bool		insertAt(1: TKey key, 2: TEntry entry, 3: i32 index);

	/**
	 * Remove an entry, return false if entry not found
	 */
	bool		remove(1: TKey key, 2: TEntry entry);

	/**
	 * Remove entry by index
	 * @param index similar to insertAt
	 */
	bool		removeAt(1: TKey key, 2: i32 index);

	/**
	 * Make an entry the first in unsorted list and return true if it existed, do nothing on sorted list return false
	 */
	bool bumpUp(1: TKey key, 2: TEntry entry);

	/**
	 * Remove entry list
	 */
	bool		removeList(1: TKey key, 2: TEntryList entryList);

	/**
	 * Clear data on given key
	 */
	void		clearData(1: TKey key);

	/**
	 * Split the list into 2, return first entry of new list, or -1 if fails
	 */
	TEntry	split(1: TKey oldKey, 2: TKey newKey, 3: i32 origLength);

	//multi key operations
	map<TKey, i32>			multiCount(1: TKeyList keyList);

	map<TKey, bool> 		multiExisted(1: TKeyList keyList, 2: TEntry entry);

	map<TKey, TEntryList> 	multiGetListAll(1: TKeyList keyList);

	map<TKey, TEntryList> 	multiGetSlice(1: TKeyList keyList, 2: i32 start, 3: i32 length);

	map<TKey, TEntryList> 	multiGetSliceReverse(1: TKeyList keyList, 2: i32 start, 3: i32 length);

	/**
	 * Put same entry to multiple keys
	 */
	map<TKey, bool>			multiPut(1: TKeyList keyList, 2: TEntry entry);

	oneway void put_ow(1: TKey key, 2: TEntry entry);
	oneway void putMultiKey_ow(1: TKeyList uidList, 2: TEntry entry);
	oneway void remove_ow(1: TKey key, 2: TEntry entry);
}
