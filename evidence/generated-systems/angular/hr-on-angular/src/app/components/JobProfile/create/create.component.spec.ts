
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateJobProfileComponent } from './create.component';
import { JobProfileService } from '../../../services/JobProfile.service';
import { Router } from '@angular/router';

describe('CreateJobProfileComponent', () => {
  let component: CreateJobProfileComponent;
  let fixture: ComponentFixture<CreateJobProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateJobProfileComponent
      ],
      providers: [
        JobProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateJobProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});