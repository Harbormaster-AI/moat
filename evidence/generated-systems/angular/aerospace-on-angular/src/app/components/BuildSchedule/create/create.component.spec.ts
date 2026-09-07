
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBuildScheduleComponent } from './create.component';
import { BuildScheduleService } from '../../../services/BuildSchedule.service';
import { Router } from '@angular/router';

describe('CreateBuildScheduleComponent', () => {
  let component: CreateBuildScheduleComponent;
  let fixture: ComponentFixture<CreateBuildScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBuildScheduleComponent
      ],
      providers: [
        BuildScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBuildScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});