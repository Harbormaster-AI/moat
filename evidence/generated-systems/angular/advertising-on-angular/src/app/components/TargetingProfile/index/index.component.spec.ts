
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTargetingProfileComponent } from './index.component';
import { TargetingProfileService } from '../../../services/TargetingProfile.service';

describe('IndexTargetingProfileComponent', () => {
  let component: IndexTargetingProfileComponent;
  let fixture: ComponentFixture<IndexTargetingProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTargetingProfileComponent
      ],
      providers: [
        TargetingProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTargetingProfileComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});