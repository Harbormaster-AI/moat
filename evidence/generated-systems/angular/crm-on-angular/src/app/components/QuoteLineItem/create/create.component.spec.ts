
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateQuoteLineItemComponent } from './create.component';
import { QuoteLineItemService } from '../../../services/QuoteLineItem.service';
import { Router } from '@angular/router';

describe('CreateQuoteLineItemComponent', () => {
  let component: CreateQuoteLineItemComponent;
  let fixture: ComponentFixture<CreateQuoteLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateQuoteLineItemComponent
      ],
      providers: [
        QuoteLineItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateQuoteLineItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});