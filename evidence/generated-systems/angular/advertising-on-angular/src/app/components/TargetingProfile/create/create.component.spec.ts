
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTargetingProfileComponent } from './create.component';
import { TargetingProfileService } from '../../../services/TargetingProfile.service';
import { Router } from '@angular/router';

describe('CreateTargetingProfileComponent', () => {
  let component: CreateTargetingProfileComponent;
  let fixture: ComponentFixture<CreateTargetingProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTargetingProfileComponent
      ],
      providers: [
        TargetingProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTargetingProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});