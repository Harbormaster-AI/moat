
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateOpportunityComponent } from './create.component';
import { OpportunityService } from '../../../services/Opportunity.service';
import { Router } from '@angular/router';

describe('CreateOpportunityComponent', () => {
  let component: CreateOpportunityComponent;
  let fixture: ComponentFixture<CreateOpportunityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateOpportunityComponent
      ],
      providers: [
        OpportunityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateOpportunityComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});