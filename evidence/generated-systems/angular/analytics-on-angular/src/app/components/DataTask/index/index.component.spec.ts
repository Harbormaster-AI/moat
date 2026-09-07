
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataTaskComponent } from './index.component';
import { DataTaskService } from '../../../services/DataTask.service';

describe('IndexDataTaskComponent', () => {
  let component: IndexDataTaskComponent;
  let fixture: ComponentFixture<IndexDataTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataTaskComponent
      ],
      providers: [
        DataTaskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataTaskComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});