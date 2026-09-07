
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMedicalSupplierComponent } from './create.component';
import { MedicalSupplierService } from '../../../services/MedicalSupplier.service';
import { Router } from '@angular/router';

describe('CreateMedicalSupplierComponent', () => {
  let component: CreateMedicalSupplierComponent;
  let fixture: ComponentFixture<CreateMedicalSupplierComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMedicalSupplierComponent
      ],
      providers: [
        MedicalSupplierService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMedicalSupplierComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});