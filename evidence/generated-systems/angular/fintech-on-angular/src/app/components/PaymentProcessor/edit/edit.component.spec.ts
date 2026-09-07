
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditPaymentProcessorComponent } from './edit.component';
import { PaymentProcessorService } from '../../../services/PaymentProcessor.service';

describe('EditPaymentProcessorComponent', () => {
  let component: EditPaymentProcessorComponent;
  let fixture: ComponentFixture<EditPaymentProcessorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditPaymentProcessorComponent
      ],
      providers: [
        PaymentProcessorService,
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

    fixture = TestBed.createComponent(EditPaymentProcessorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});