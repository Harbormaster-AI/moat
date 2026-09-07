
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateChargebackComponent } from './create.component';
import { ChargebackService } from '../../../services/Chargeback.service';
import { Router } from '@angular/router';

describe('CreateChargebackComponent', () => {
  let component: CreateChargebackComponent;
  let fixture: ComponentFixture<CreateChargebackComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateChargebackComponent
      ],
      providers: [
        ChargebackService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateChargebackComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});