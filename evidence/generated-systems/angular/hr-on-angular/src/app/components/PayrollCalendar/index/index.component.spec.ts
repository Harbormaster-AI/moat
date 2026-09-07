
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPayrollCalendarComponent } from './index.component';
import { PayrollCalendarService } from '../../../services/PayrollCalendar.service';

describe('IndexPayrollCalendarComponent', () => {
  let component: IndexPayrollCalendarComponent;
  let fixture: ComponentFixture<IndexPayrollCalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPayrollCalendarComponent
      ],
      providers: [
        PayrollCalendarService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPayrollCalendarComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});