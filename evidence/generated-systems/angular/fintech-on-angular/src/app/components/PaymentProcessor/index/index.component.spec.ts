
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPaymentProcessorComponent } from './index.component';
import { PaymentProcessorService } from '../../../services/PaymentProcessor.service';

describe('IndexPaymentProcessorComponent', () => {
  let component: IndexPaymentProcessorComponent;
  let fixture: ComponentFixture<IndexPaymentProcessorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPaymentProcessorComponent
      ],
      providers: [
        PaymentProcessorService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPaymentProcessorComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});