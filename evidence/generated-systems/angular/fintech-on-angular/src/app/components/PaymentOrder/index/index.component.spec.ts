
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPaymentOrderComponent } from './index.component';
import { PaymentOrderService } from '../../../services/PaymentOrder.service';

describe('IndexPaymentOrderComponent', () => {
  let component: IndexPaymentOrderComponent;
  let fixture: ComponentFixture<IndexPaymentOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPaymentOrderComponent
      ],
      providers: [
        PaymentOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPaymentOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});