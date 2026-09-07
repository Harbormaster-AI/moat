
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditOpportunityLineItemComponent } from './edit.component';
import { OpportunityLineItemService } from '../../../services/OpportunityLineItem.service';

describe('EditOpportunityLineItemComponent', () => {
  let component: EditOpportunityLineItemComponent;
  let fixture: ComponentFixture<EditOpportunityLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditOpportunityLineItemComponent
      ],
      providers: [
        OpportunityLineItemService,
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

    fixture = TestBed.createComponent(EditOpportunityLineItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});