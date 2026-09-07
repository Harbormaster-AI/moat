
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditMedicalDeviceComponent } from './edit.component';
import { MedicalDeviceService } from '../../../services/MedicalDevice.service';

describe('EditMedicalDeviceComponent', () => {
  let component: EditMedicalDeviceComponent;
  let fixture: ComponentFixture<EditMedicalDeviceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditMedicalDeviceComponent
      ],
      providers: [
        MedicalDeviceService,
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

    fixture = TestBed.createComponent(EditMedicalDeviceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});