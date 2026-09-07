
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPaymentMethodComponent } from './index.component';
import { PaymentMethodService } from '../../../services/PaymentMethod.service';

describe('IndexPaymentMethodComponent', () => {
  let component: IndexPaymentMethodComponent;
  let fixture: ComponentFixture<IndexPaymentMethodComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPaymentMethodComponent
      ],
      providers: [
        PaymentMethodService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPaymentMethodComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});