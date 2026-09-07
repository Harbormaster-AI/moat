
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexMaintenanceAppointmentComponent } from './index.component';
import { MaintenanceAppointmentService } from '../../../services/MaintenanceAppointment.service';

describe('IndexMaintenanceAppointmentComponent', () => {
  let component: IndexMaintenanceAppointmentComponent;
  let fixture: ComponentFixture<IndexMaintenanceAppointmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexMaintenanceAppointmentComponent
      ],
      providers: [
        MaintenanceAppointmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexMaintenanceAppointmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});