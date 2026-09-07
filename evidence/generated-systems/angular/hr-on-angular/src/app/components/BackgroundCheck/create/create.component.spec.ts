
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBackgroundCheckComponent } from './create.component';
import { BackgroundCheckService } from '../../../services/BackgroundCheck.service';
import { Router } from '@angular/router';

describe('CreateBackgroundCheckComponent', () => {
  let component: CreateBackgroundCheckComponent;
  let fixture: ComponentFixture<CreateBackgroundCheckComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBackgroundCheckComponent
      ],
      providers: [
        BackgroundCheckService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBackgroundCheckComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});