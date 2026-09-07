
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateOpportunityLineItemComponent } from './create.component';
import { OpportunityLineItemService } from '../../../services/OpportunityLineItem.service';
import { Router } from '@angular/router';

describe('CreateOpportunityLineItemComponent', () => {
  let component: CreateOpportunityLineItemComponent;
  let fixture: ComponentFixture<CreateOpportunityLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateOpportunityLineItemComponent
      ],
      providers: [
        OpportunityLineItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateOpportunityLineItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});