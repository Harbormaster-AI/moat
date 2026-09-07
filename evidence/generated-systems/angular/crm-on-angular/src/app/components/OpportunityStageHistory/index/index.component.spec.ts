
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexOpportunityStageHistoryComponent } from './index.component';
import { OpportunityStageHistoryService } from '../../../services/OpportunityStageHistory.service';

describe('IndexOpportunityStageHistoryComponent', () => {
  let component: IndexOpportunityStageHistoryComponent;
  let fixture: ComponentFixture<IndexOpportunityStageHistoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexOpportunityStageHistoryComponent
      ],
      providers: [
        OpportunityStageHistoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexOpportunityStageHistoryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});