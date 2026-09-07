
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreatePayrollCalendarComponent } from './create.component';
import { PayrollCalendarService } from '../../../services/PayrollCalendar.service';
import { Router } from '@angular/router';

describe('CreatePayrollCalendarComponent', () => {
  let component: CreatePayrollCalendarComponent;
  let fixture: ComponentFixture<CreatePayrollCalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreatePayrollCalendarComponent
      ],
      providers: [
        PayrollCalendarService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreatePayrollCalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});