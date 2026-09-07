
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTimesheetComponent } from './create.component';
import { TimesheetService } from '../../../services/Timesheet.service';
import { Router } from '@angular/router';

describe('CreateTimesheetComponent', () => {
  let component: CreateTimesheetComponent;
  let fixture: ComponentFixture<CreateTimesheetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTimesheetComponent
      ],
      providers: [
        TimesheetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTimesheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});