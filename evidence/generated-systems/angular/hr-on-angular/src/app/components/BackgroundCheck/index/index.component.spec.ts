
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBackgroundCheckComponent } from './index.component';
import { BackgroundCheckService } from '../../../services/BackgroundCheck.service';

describe('IndexBackgroundCheckComponent', () => {
  let component: IndexBackgroundCheckComponent;
  let fixture: ComponentFixture<IndexBackgroundCheckComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBackgroundCheckComponent
      ],
      providers: [
        BackgroundCheckService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBackgroundCheckComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});