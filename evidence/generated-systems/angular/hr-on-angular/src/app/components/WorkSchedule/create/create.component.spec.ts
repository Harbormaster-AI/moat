
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateWorkScheduleComponent } from './create.component';
import { WorkScheduleService } from '../../../services/WorkSchedule.service';
import { Router } from '@angular/router';

describe('CreateWorkScheduleComponent', () => {
  let component: CreateWorkScheduleComponent;
  let fixture: ComponentFixture<CreateWorkScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateWorkScheduleComponent
      ],
      providers: [
        WorkScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateWorkScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});