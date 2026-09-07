
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditOpportunityStageHistoryComponent } from './edit.component';
import { OpportunityStageHistoryService } from '../../../services/OpportunityStageHistory.service';

describe('EditOpportunityStageHistoryComponent', () => {
  let component: EditOpportunityStageHistoryComponent;
  let fixture: ComponentFixture<EditOpportunityStageHistoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditOpportunityStageHistoryComponent
      ],
      providers: [
        OpportunityStageHistoryService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditOpportunityStageHistoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});