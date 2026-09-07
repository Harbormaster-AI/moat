
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRetentionScheduleComponent } from './create.component';
import { RetentionScheduleService } from '../../../services/RetentionSchedule.service';
import { Router } from '@angular/router';

describe('CreateRetentionScheduleComponent', () => {
  let component: CreateRetentionScheduleComponent;
  let fixture: ComponentFixture<CreateRetentionScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRetentionScheduleComponent
      ],
      providers: [
        RetentionScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRetentionScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});