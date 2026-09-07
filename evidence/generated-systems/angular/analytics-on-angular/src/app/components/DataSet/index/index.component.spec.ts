
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataSetComponent } from './index.component';
import { DataSetService } from '../../../services/DataSet.service';

describe('IndexDataSetComponent', () => {
  let component: IndexDataSetComponent;
  let fixture: ComponentFixture<IndexDataSetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataSetComponent
      ],
      providers: [
        DataSetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataSetComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});