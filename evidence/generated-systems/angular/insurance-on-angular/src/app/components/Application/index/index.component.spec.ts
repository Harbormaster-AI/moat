
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexApplicationComponent } from './index.component';
import { ApplicationService } from '../../../services/Application.service';

describe('IndexApplicationComponent', () => {
  let component: IndexApplicationComponent;
  let fixture: ComponentFixture<IndexApplicationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexApplicationComponent
      ],
      providers: [
        ApplicationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexApplicationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});