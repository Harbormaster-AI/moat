
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMaintenanceAppointmentComponent } from './create.component';
import { MaintenanceAppointmentService } from '../../../services/MaintenanceAppointment.service';
import { Router } from '@angular/router';

describe('CreateMaintenanceAppointmentComponent', () => {
  let component: CreateMaintenanceAppointmentComponent;
  let fixture: ComponentFixture<CreateMaintenanceAppointmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMaintenanceAppointmentComponent
      ],
      providers: [
        MaintenanceAppointmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMaintenanceAppointmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});