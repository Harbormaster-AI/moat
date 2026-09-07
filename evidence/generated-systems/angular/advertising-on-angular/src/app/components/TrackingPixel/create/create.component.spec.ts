
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTrackingPixelComponent } from './create.component';
import { TrackingPixelService } from '../../../services/TrackingPixel.service';
import { Router } from '@angular/router';

describe('CreateTrackingPixelComponent', () => {
  let component: CreateTrackingPixelComponent;
  let fixture: ComponentFixture<CreateTrackingPixelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTrackingPixelComponent
      ],
      providers: [
        TrackingPixelService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTrackingPixelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});