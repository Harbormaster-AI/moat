
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMedicationDispenseComponent } from './create.component';
import { MedicationDispenseService } from '../../../services/MedicationDispense.service';
import { Router } from '@angular/router';

describe('CreateMedicationDispenseComponent', () => {
  let component: CreateMedicationDispenseComponent;
  let fixture: ComponentFixture<CreateMedicationDispenseComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMedicationDispenseComponent
      ],
      providers: [
        MedicationDispenseService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMedicationDispenseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});