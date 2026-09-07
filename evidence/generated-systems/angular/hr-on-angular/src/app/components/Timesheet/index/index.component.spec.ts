
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTimesheetComponent } from './index.component';
import { TimesheetService } from '../../../services/Timesheet.service';

describe('IndexTimesheetComponent', () => {
  let component: IndexTimesheetComponent;
  let fixture: ComponentFixture<IndexTimesheetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTimesheetComponent
      ],
      providers: [
        TimesheetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTimesheetComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});