import { createSelector, createFeatureSelector } from '@ngrx/store';
import { EnvironmentEntityState, environmentEntityAdapter } from './environment.reducer';

export const environmentState = createFeatureSelector<EnvironmentEntityState>('environments');
export const {
  selectAll: allEnvironments,
  selectEntities: environmentEntities
} = environmentEntityAdapter.getSelectors(environmentState);

export const environmentStatus = createSelector(
  environmentState,
  (state) => state.environmentsStatus
);

export const getAllStatus = createSelector(
  environmentState,
  (state) => state.getAllStatus
);

export const deleteStatus = createSelector(
  environmentState,
  (state) => state.deleteStatus
);

export const environmentList = createSelector(
  environmentState,
  (state) => state.environmentList
);
