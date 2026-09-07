
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditMaintenanceAppointmentComponent } from './edit.component';
import { MaintenanceAppointmentService } from '../../../services/MaintenanceAppointment.service';

describe('EditMaintenanceAppointmentComponent', () => {
  let component: EditMaintenanceAppointmentComponent;
  let fixture: ComponentFixture<EditMaintenanceAppointmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditMaintenanceAppointmentComponent
      ],
      providers: [
        MaintenanceAppointmentService,
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

    fixture = TestBed.createComponent(EditMaintenanceAppointmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});