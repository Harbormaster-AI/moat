
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAppointmentComponent } from './index.component';
import { AppointmentService } from '../../../services/Appointment.service';

describe('IndexAppointmentComponent', () => {
  let component: IndexAppointmentComponent;
  let fixture: ComponentFixture<IndexAppointmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAppointmentComponent
      ],
      providers: [
        AppointmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAppointmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});