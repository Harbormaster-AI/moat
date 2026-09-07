
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateOpportunityStageHistoryComponent } from './create.component';
import { OpportunityStageHistoryService } from '../../../services/OpportunityStageHistory.service';
import { Router } from '@angular/router';

describe('CreateOpportunityStageHistoryComponent', () => {
  let component: CreateOpportunityStageHistoryComponent;
  let fixture: ComponentFixture<CreateOpportunityStageHistoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateOpportunityStageHistoryComponent
      ],
      providers: [
        OpportunityStageHistoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateOpportunityStageHistoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});