
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexQuoteLineItemComponent } from './index.component';
import { QuoteLineItemService } from '../../../services/QuoteLineItem.service';

describe('IndexQuoteLineItemComponent', () => {
  let component: IndexQuoteLineItemComponent;
  let fixture: ComponentFixture<IndexQuoteLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexQuoteLineItemComponent
      ],
      providers: [
        QuoteLineItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexQuoteLineItemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});