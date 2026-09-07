
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMedicalDeviceComponent } from './create.component';
import { MedicalDeviceService } from '../../../services/MedicalDevice.service';
import { Router } from '@angular/router';

describe('CreateMedicalDeviceComponent', () => {
  let component: CreateMedicalDeviceComponent;
  let fixture: ComponentFixture<CreateMedicalDeviceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMedicalDeviceComponent
      ],
      providers: [
        MedicalDeviceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMedicalDeviceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});