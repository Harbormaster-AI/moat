
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMedicationOrderComponent } from './create.component';
import { MedicationOrderService } from '../../../services/MedicationOrder.service';
import { Router } from '@angular/router';

describe('CreateMedicationOrderComponent', () => {
  let component: CreateMedicationOrderComponent;
  let fixture: ComponentFixture<CreateMedicationOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMedicationOrderComponent
      ],
      providers: [
        MedicationOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMedicationOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});