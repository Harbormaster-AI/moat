
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateScheduleExceptionComponent } from './create.component';
import { ScheduleExceptionService } from '../../../services/ScheduleException.service';
import { Router } from '@angular/router';

describe('CreateScheduleExceptionComponent', () => {
  let component: CreateScheduleExceptionComponent;
  let fixture: ComponentFixture<CreateScheduleExceptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateScheduleExceptionComponent
      ],
      providers: [
        ScheduleExceptionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateScheduleExceptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});