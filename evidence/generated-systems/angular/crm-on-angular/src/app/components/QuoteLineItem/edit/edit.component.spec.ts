
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditQuoteLineItemComponent } from './edit.component';
import { QuoteLineItemService } from '../../../services/QuoteLineItem.service';

describe('EditQuoteLineItemComponent', () => {
  let component: EditQuoteLineItemComponent;
  let fixture: ComponentFixture<EditQuoteLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditQuoteLineItemComponent
      ],
      providers: [
        QuoteLineItemService,
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

    fixture = TestBed.createComponent(EditQuoteLineItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});