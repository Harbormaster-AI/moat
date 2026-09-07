
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexCareTaskComponent } from './index.component';
import { CareTaskService } from '../../../services/CareTask.service';

describe('IndexCareTaskComponent', () => {
  let component: IndexCareTaskComponent;
  let fixture: ComponentFixture<IndexCareTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexCareTaskComponent
      ],
      providers: [
        CareTaskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexCareTaskComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});